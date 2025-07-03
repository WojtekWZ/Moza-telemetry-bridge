package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var replacer = strings.NewReplacer(", %", "",
	", km/h", "",
	", m/s", "",
	", m", "",
	", deg/s", "",
	", deg", "",
	", kgs", "",
	", kg", "",
	", hp", "",
	", atm", "",
	", C", "",
)

var totalShells = 1900

func main() {
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Timeout: 100 * time.Millisecond,
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn, client)
	}

}

func handleConnection(conn net.Conn, client *http.Client) {
	defer conn.Close()

	wg := new(sync.WaitGroup)
	for {
		if totalShells <= 0 {
			totalShells = 30000
		}
		var telem DCSResponse
		wg.Add(2)
		go getWtState(client, &telem, wg)
		go getWtIndicators(client, &telem, wg)
		wg.Wait()
		_, err := conn.Write([]byte(telem.toString()))
		if err != nil {
			fmt.Println(err)
		}
	}

}

func getWtState(client *http.Client, telem *DCSResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		resp, err := client.Get("http://127.0.0.1:8111/state")
		if err != nil || resp.StatusCode != 200 {
			fmt.Printf("err conn: %v\n", err)
			if !os.IsTimeout(err) {
				time.Sleep(500 * time.Millisecond)
			}
			continue
		}
		if resp.StatusCode != 200 {
			fmt.Printf("status: %v\n", resp.StatusCode)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if resp.Body != nil {
			state, err := readStatusResponse(resp)
			if err != nil {
				continue
			}
			wtStateToDCSResponse(state, telem)
			return
		}
	}
}

func getWtIndicators(client *http.Client, telem *DCSResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		resp, err := client.Get("http://127.0.0.1:8111/indicators")
		if err != nil || resp.StatusCode != 200 {
			fmt.Printf("err conn: %v\n", err)
			if !os.IsTimeout(err) {
				time.Sleep(500 * time.Millisecond)
			}
			continue
		}
		if resp.StatusCode != 200 {
			fmt.Printf("status: %v\n", resp.StatusCode)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if resp.Body != nil {
			indicators, err := readIndicatorsResponse(resp)
			if err != nil {
				continue
			}
			if indicators.Weapon1 > 0 || indicators.Weapon2 > 0 {
				totalShells -= 2
			}
			telem.cannon_shells = totalShells
			telem.bank = indicators.Bank
			return
		}
	}
}

func readStatusResponse(response *http.Response) (wtState *WTState, err error) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	var state WTState
	err = json.Unmarshal([]byte(replacer.Replace(string(body))), &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func readIndicatorsResponse(response *http.Response) (wtState *WTIndicators, err error) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	var indicators WTIndicators
	err = json.Unmarshal(body, &indicators)
	if err != nil {
		return nil, err
	}
	return &indicators, nil
}

func wtStateToDCSResponse(wtState *WTState, telem *DCSResponse) {
	telem.engine_rpm_left = float64(wtState.Throttle1)
	telem.engine_rpm_right = float64(wtState.Throttle2)
	telem.left_gear = wtState.Gear / 100
	telem.nose_gear = wtState.Gear / 100
	telem.right_gear = wtState.Gear / 100
	telem.acc_x = 0
	telem.acc_y = wtState.Ny
	telem.acc_z = 0
	telem.vector_velocity_x = 0
	telem.vector_velocity_y = wtState.VyMS
	telem.vector_velocity_z = float64(wtState.IASKmH) / 3.6
	telem.tas = float64(wtState.TASKmH)
	telem.ias = float64(wtState.IASKmH)
	telem.vertical_velocity_speed = wtState.VyMS
	telem.aoa = wtState.AoADeg
	telem.pitch = wtState.Pitch1Deg
	telem.aos = wtState.AoSDeg
	telem.flap_pos = wtState.Flaps / 100
	telem.gear_value = wtState.Gear / 100
	telem.speedbrake_value = wtState.Airbrake / 100
	telem.afterburner_1 = 0
	telem.afterburner_2 = 0
	telem.mach = wtState.M
	telem.h_above_sea_level = float64(wtState.HM)
}
