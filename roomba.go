// Copyright (c) 2020 Peter Hagelund
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package irobot

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/peterhagelund/go-serial"
)

// OpCode is the OI command op-code type.
type OpCode byte

const (
	// OpCodeStart stats the OI.
	OpCodeStart OpCode = iota + 128
	// OpCodeBaud changes the baud rate of the OI.
	OpCodeBaud
	// OpCodeControl ...
	OpCodeControl
	// OpCodeSafe puts the OI in safe mode.
	OpCodeSafe
	// OpCodeFull puts the OI in full mode.
	OpCodeFull
	// OpCodePower powers off the OI.
	OpCodePower
	// OpCodeSpot starts a spot cleaning.
	OpCodeSpot
	// OpCodeClean starts a default cleaning.
	OpCodeClean
	// OpCodeMax starts a max cleaning.
	OpCodeMax
	// OpCodeDrive controls the drive wheels.
	OpCodeDrive
	// OpCodeMotors controls the brush and vacuum motors.
	OpCodeMotors
	// OpCodeLEDs controls the LEDs.
	OpCodeLEDs
	// OpCodeSong defines a song.
	OpCodeSong
	// OpCodePlay plays a song.
	OpCodePlay
	// OpCodeSensors requests a sensor packet.
	OpCodeSensors
	// OpCodeSeekDock instructs the Roomba tp seek its dock.
	OpCodeSeekDock
	// OpCodeMotorsPWM controls the brush and vacuum motors using pulse width modulation (PWM).
	OpCodeMotorsPWM
	// OpCodeDriveDirect controls the wheels using direct drive.
	OpCodeDriveDirect
	// OpCodeDrivePWM controls the wheel using pulse width modulation (PWM).
	OpCodeDrivePWM
	// Missing op-code.
	_
	// OpCodeStream instructs the Roomba to start streaming packets every 15 ms.
	OpCodeStream
	// OpCodeQueryList queries (once) a list of sensor packets.
	OpCodeQueryList
	// OpCodePauseResumeStream pauses and resumes the streaming of packets.
	OpCodePauseResumeStream
	// Missing op-codes.
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	// OpCodeSchedulingLEDs ...
	OpCodeSchedulingLEDs
	// OpCodeDigitLEDsRaw ...
	OpCodeDigitLEDsRaw
	// OpCodeDigitLEDsASCII ...
	OpCodeDigitLEDsASCII
	// OpCodeButtons ...
	OpCodeButtons
	// Missing op-code.
	_
	// OpCodeSchedule ...
	OpCodeSchedule
	// OpCodeSetDayTime ...
	OpCodeSetDayTime
)

// BaudRate is the OI baud rate type.
type BaudRate byte

const (
	// BaudRate300 is the 300 bps baud rate.
	BaudRate300 BaudRate = iota
	// BaudRate600 is the 600 bps baud rate.
	BaudRate600
	// BaudRate1200 is the 1200 bps baud rate.
	BaudRate1200
	// BaudRate2400 is the 2400 bps baud rate.
	BaudRate2400
	// BaudRate4800 is the 4800 bps baud rate.
	BaudRate4800
	// BaudRate9600 is the 9600 bps baud rate.
	BaudRate9600
	// BaudRate14400 is the 14400 bps baud rate.
	BaudRate14400
	// BaudRate19200 is the 19200 bps baud rate.
	BaudRate19200
	// BaudRate28800 is the 28800 bps baud rate.
	BaudRate28800
	// BaudRate38400 is the 38300 bps baud rate.
	BaudRate38400
	// BaudRate57600 is the 57600 bps baud rate.
	BaudRate57600
	// BaudRate115200 is the 115200 bps baud rate.
	BaudRate115200
)

// MainBrush is the main brush motor control type.
type MainBrush byte

const (
	// MainBrushOff turns the main brush off.
	MainBrushOff MainBrush = iota
	// MainBrushInward turns the main brush on rotating in an inward direction.
	MainBrushInward
	// MainBrushOutward turns the main brush on rotating in an outward direction.
	MainBrushOutward
)

// SideBrush is the side brush motor control type.
type SideBrush byte

const (
	// SideBrushOff turns the side brush off.
	SideBrushOff SideBrush = iota
	// SideBrushClockwise turns the side brush on rotating in a clockwise direction.
	SideBrushClockwise
	// SideBrushCounterClockwise turns the side brush on rotating in a counter clockwise direction.
	SideBrushCounterClockwise
)

// Vacuum is the vacuum motor control type.
type Vacuum byte

const (
	// VacuumOff turns the vacuum motor off.
	VacuumOff Vacuum = iota
	// VacuumOn turns the vacuum motor on.
	VacuumOn
)

// ChargingState is the charging state type.
type ChargingState byte

const (
	// ChargingStateNotCharging indicates the battery is not charging.
	ChargingStateNotCharging ChargingState = iota
	// ChargingStateReconditioningCharging indicates the battery is being reconditioning charged.
	ChargingStateReconditioningCharging
	// ChargingStateFullCharging indicates the battery is being full charged.
	ChargingStateFullCharging
	// ChargingStateTrickleCharging indicates the battery is being trickle charged.
	ChargingStateTrickleCharging
	// ChargingStateWaiting indicates the Roomba is waiting to charge the battery.
	ChargingStateWaiting
	// ChargingStateChargingFaultCondition indicates a fault condition while charging.
	ChargingStateChargingFaultCondition
	// ChargingStateUnknown indicates an unknown charging state.
	ChargingStateUnknown
)

// Mode is the OI mode type.
type Mode byte

const (
	// ModeOff indicates the OI is in "off" mode.
	ModeOff Mode = iota
	// ModePassive indicates the OI is in "passive" mode.
	ModePassive
	// ModeSafe indicates the OI is in "safe" mode.
	ModeSafe
	// ModeFull indicates the OI is in "full" mode.
	ModeFull
	// ModeUnknown indicates the OI is in "unknown" mode.
	ModeUnknown
)

// Note is an encapsulation of a single Roomba song note.
type Note struct {
	// Number is the note number (31 to 127)
	Number uint8
	// Duration is the note duration in 1/64th of a second (0 to 255)
	Duration uint8
}

// Roomba defines the required behavior of an iRobot Roomba vacuum cleaner OI.
type Roomba interface {
	// Start starts the OI.
	Start() error
	// SetBaudRate sets the baud rate of the OI.
	SetBaudRate(baudRate BaudRate) error
	// Safe puts the OI in safe mode.
	Safe() error
	// Full puts the OI in full mode.
	Full() error
	// Power powers off the OI.
	Power() error
	// Spot starts a spot cleaning.
	Spot() error
	// Clean starts a default cleaning.
	Clean() error
	// Max starts a max cleaning.
	Max() error
	// Drive controls the drive wheels.
	Drive(velocity int16, radius int16) error
	// Motors controls the brush and vacuum motors.
	Motors(mainBrush MainBrush, sideBrush SideBrush, vacuum Vacuum) error
	// LEDs controls the various Roomba LEDs.
	LEDs(color uint8, intensity uint8, debris bool, spot bool, dock bool, checkRobot bool) error
	// SetSong defines a song.
	SetSong(number uint8, note []*Note) error
	// Play instructs the Roomba to play the specified song.
	Play(number uint8) error
	// Sensors requests a sensor value packet from the Roomba.
	Sensors(id int) (Packet, error)
	// UpdateSensors requests an update for the specified sensor packet.
	UpdateSensors(packet Packet) error
	// MotorsPWM controls the brush and vacuum motors directlyt using Pulse Width Modulation (PWM).
	MotorsPWM(mainBrushPWM int8, sideBrushPWM int8, vacuumPWM uint8) error
	// DriveDirect controls the drive wheels directly by setting a velocity for each.
	DriveDirect(leftVelocity int16, rightVelocity int16) error
	// DrivePWM controls the drive wheels directly using Pulse Width Modulation (PWM).
	DrivePWM(leftWheelPWM int16, rightWheelPWM int16) error
}

type roomba struct {
	conn      net.Conn
	lastWrite time.Time
}

var baudRateMap = map[BaudRate]serial.BaudRate{
	BaudRate300:    serial.BaudRate300,
	BaudRate600:    serial.BaudRate600,
	BaudRate1200:   serial.BaudRate1200,
	BaudRate2400:   serial.BaudRate2400,
	BaudRate4800:   serial.BaudRate4800,
	BaudRate9600:   serial.BaudRate9600,
	BaudRate14400:  serial.BaudRate14400,
	BaudRate19200:  serial.BaudRate19200,
	BaudRate28800:  serial.BaudRate28800,
	BaudRate38400:  serial.BaudRate38400,
	BaudRate57600:  serial.BaudRate57600,
	BaudRate115200: serial.BaudRate115200,
}

// NewNote creates and returns a new Note with the specified number and duration.
func NewNote(number uint8, duration time.Duration) (*Note, error) {
	if number < 31 || number > 127 {
		return nil, fmt.Errorf("invalid number (%d)", number)
	}
	d := duration / time.Millisecond * 64 / 1000
	if d > 255 {
		return nil, fmt.Errorf("invalid duration (%d)", duration)
	}
	return &Note{
		Number:   number,
		Duration: uint8(d),
	}, nil
}

// NewRoomba creates and returns a new Roomba instance for the specified connection.
func NewRoomba(conn net.Conn) (Roomba, error) {
	if conn == nil {
		return nil, errors.New("nil connection supplied")
	}
	return &roomba{
		conn:      conn,
		lastWrite: time.Now(),
	}, nil
}

func (r *roomba) Start() error {
	data := make([]byte, 1)
	data[0] = byte(OpCodeStart)
	return r.write(data)
}

func (r *roomba) SetBaudRate(baudRate BaudRate) error {
	conn, ok := r.conn.(serial.PortConn)
	if !ok {
		return errors.New("connection not serial")
	}
	b, ok := baudRateMap[baudRate]
	if !ok {
		return errors.New("invalid baud rate")
	}
	data := make([]byte, 12)
	data[0] = byte(OpCodeBaud)
	data[1] = byte(baudRate)
	if err := r.write(data); err != nil {
		return err
	}
	// Allow Roomba to switch baud rate
	time.Sleep(100 * time.Millisecond)
	return conn.Port().SetBaudRate(b)
}

func (r *roomba) Safe() error {
	data := make([]byte, 1)
	data[0] = byte(OpCodeSafe)
	return r.write(data)
}

func (r *roomba) Full() error {
	data := make([]byte, 1)
	data[0] = byte(OpCodeFull)
	return r.write(data)
}

func (r *roomba) Power() error {
	data := make([]byte, 1)
	data[0] = byte(OpCodePower)
	return r.write(data)
}

func (r *roomba) Spot() error {
	data := make([]byte, 1)
	data[0] = byte(OpCodeSpot)
	return r.write(data)
}

func (r *roomba) Clean() error {
	data := make([]byte, 1)
	data[0] = byte(OpCodeClean)
	return r.write(data)
}

func (r *roomba) Max() error {
	data := make([]byte, 1)
	data[0] = byte(OpCodeMax)
	return r.write(data)
}

func (r *roomba) Drive(velocity int16, radius int16) error {
	if velocity < -500 || velocity > 500 {
		return fmt.Errorf("invalid velocity (%d)", velocity)
	}
	if radius < -2000 || radius > 2000 {
		return fmt.Errorf("invalid radius (%d)", radius)
	}
	data := make([]byte, 5)
	data[0] = byte(OpCodeDrive)
	data[1], data[2] = int16ToBytes(velocity)
	data[3], data[4] = int16ToBytes(radius)
	return r.write(data)
}

func (r *roomba) Motors(mainBrush MainBrush, sideBrush SideBrush, vacuum Vacuum) error {
	data := make([]byte, 2)
	data[0] = byte(OpCodeMotors)
	var bits byte = 0b00000000
	switch mainBrush {
	case MainBrushOff:
		break
	case MainBrushInward:
		bits |= 0b00000100
	case MainBrushOutward:
		bits |= 0b00010100
	}
	switch sideBrush {
	case SideBrushOff:
		break
	case SideBrushCounterClockwise:
		bits |= 0b00000001
	case SideBrushClockwise:
		bits |= 0b00001001
	}
	switch vacuum {
	case VacuumOff:
		break
	case VacuumOn:
		bits |= 0b00000010
	}
	data[1] = bits
	return r.write(data)
}

func (r *roomba) LEDs(color uint8, intensity uint8, debris bool, spot bool, dock bool, checkRobot bool) error {
	data := make([]byte, 4)
	data[0] = byte(OpCodeLEDs)
	var bits byte = 0b00000000
	if debris {
		bits |= 0b00000001
	}
	if spot {
		bits |= 0b00000010
	}
	if dock {
		bits |= 0b00000100
	}
	if checkRobot {
		bits |= 0b00001000
	}
	data[1] = bits
	data[2] = color
	data[3] = intensity
	return r.write(data)
}

func (r *roomba) SetSong(number uint8, notes []*Note) error {
	if number > 4 {
		return fmt.Errorf("invalid song number (%d)", number)
	}
	noteCount := len(notes)
	if noteCount < 1 || noteCount > 16 {
		return fmt.Errorf("invalid number of notes (%d)", noteCount)
	}
	for i, note := range notes {
		if note.Number < 31 || note.Number > 127 {
			return fmt.Errorf("note at index %d has invalid number", i)
		}
	}
	data := make([]byte, 3+2*noteCount)
	data[0] = byte(OpCodeSong)
	data[1] = number
	data[2] = byte(len(notes))
	for i, note := range notes {
		data[3+i*2+0] = note.Number
		data[3+i*2+1] = note.Duration
	}
	return r.write(data)
}

func (r *roomba) Play(number uint8) error {
	if number > 4 {
		return fmt.Errorf("invalid song number (%d)", number)
	}
	data := make([]byte, 2)
	data[0] = byte(OpCodePlay)
	data[1] = number
	return r.write(data)
}

func (r *roomba) Sensors(id int) (Packet, error) {
	packet, err := NewPacket(id)
	if err != nil {
		return nil, err
	}
	err = r.UpdateSensors(packet)
	if err != nil {
		return nil, err
	}
	return packet, nil
}

func (r *roomba) UpdateSensors(packet Packet) error {
	// TODO
	// send sensor op-code
	// read sensor data
	// extract
	return nil
}

func (r *roomba) MotorsPWM(mainBrushPWM int8, sideBrushPWM int8, vacuumPWM uint8) error {
	if mainBrushPWM < -127 {
		return fmt.Errorf("invalid main brush PWM value (%d)", mainBrushPWM)
	}
	if sideBrushPWM < -127 {
		return fmt.Errorf("invalid side brush PWM value (%d)", sideBrushPWM)
	}
	if vacuumPWM > 127 {
		return fmt.Errorf("invalid vacuum PWM value (%d)", vacuumPWM)
	}
	data := make([]byte, 4)
	data[0] = byte(OpCodeMotorsPWM)
	data[1] = byte(mainBrushPWM)
	data[2] = byte(sideBrushPWM)
	data[3] = vacuumPWM
	return r.write(data)
}

func (r *roomba) DriveDirect(leftVelocity int16, rightVelocity int16) error {
	if leftVelocity < -500 || leftVelocity > 500 {
		return fmt.Errorf("invalid left velocity (%d)", leftVelocity)
	}
	if rightVelocity < -500 || rightVelocity > 500 {
		return fmt.Errorf("invalid right velocity (%d)", rightVelocity)
	}
	data := make([]byte, 5)
	data[0] = byte(OpCodeDriveDirect)
	data[1], data[2] = int16ToBytes(rightVelocity)
	data[3], data[4] = int16ToBytes(leftVelocity)
	return r.write(data)
}

func (r *roomba) DrivePWM(leftWheelPWM int16, rightWheelPWM int16) error {
	if leftWheelPWM < -255 || leftWheelPWM > 255 {
		return fmt.Errorf("invalid left wheel PWM value (%d)", leftWheelPWM)
	}
	if rightWheelPWM < -255 || rightWheelPWM > 255 {
		return fmt.Errorf("invalid right wheel PWM value (%d)", rightWheelPWM)
	}
	data := make([]byte, 5)
	data[0] = byte(OpCodeDrivePWM)
	data[1], data[2] = int16ToBytes(rightWheelPWM)
	data[3], data[4] = int16ToBytes(leftWheelPWM)
	return r.write(data)
}

func (r *roomba) write(data []byte) error {
	duration := time.Now().Sub(r.lastWrite)
	if duration < 50*time.Millisecond {
		time.Sleep(duration)
	}
	n, err := r.conn.Write(data)
	if err != nil {
		return err
	}
	r.lastWrite = time.Now()
	if n < len(data) {
		return fmt.Errorf("incomplete write (%d)", n)
	}
	return nil
}
