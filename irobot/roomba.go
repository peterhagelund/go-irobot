// Copyright (c) 2020-2024 Peter Hagelund
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
	"io"
	"sync"
	"time"

	"go.bug.st/serial"
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
	// OpCodeSeekDock instructs the Roomba to seek its dock.
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

const (
	// ProcessTimeDefault is the default processing time for most op-codes.
	ProcessTimeStart time.Duration = 500 * time.Millisecond
	// ProcessTimeBaudRate is th processing time for changing baud rate.
	ProcessTimeModeChange = 50 * time.Millisecond
	// ProcessTimeModeChange is the processing time for the op-codes that change mode.
	ProcessTimeCommand = 25 * time.Millisecond
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
	io.ReadWriteCloser
	// Start starts the OI.
	Start() error
	// SetBaudRate sets the baud rate of the OI.
	SetBaudRate(baudRate BaudRate) error
	// Control enables control.
	Control() error
	// Safe puts the OI in safe mode.
	Safe() error
	// Full puts the OI in full mode.
	Full() error
	// Power powers off the OI.
	Power() error
	// SetMode instructs teh Roomba to enter the specified mode.
	SetMode(mode Mode) error
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
	// SeekDock instructs the Roomba to seek its dock.
	SeekDock() error
	// MotorsPWM controls the brush and vacuum motors directlyt using Pulse Width Modulation (PWM).
	MotorsPWM(mainBrushPWM int8, sideBrushPWM int8, vacuumPWM uint8) error
	// DriveDirect controls the drive wheels directly by setting a velocity for each.
	DriveDirect(leftVelocity int16, rightVelocity int16) error
	// DrivePWM controls the drive wheels directly using Pulse Width Modulation (PWM).
	DrivePWM(leftWheelPWM int16, rightWheelPWM int16) error
	// Stream instructs the Roomba to stream sensor packets every 15 ms.
	Stream(ids ...int) (int, error)
	// QueryList requests a list of sensor packets from the Roomba.
	QueryList(ids ...int) ([]Packet, error)
	// UpdateList updates the list of sensor packets.
	UpdateList(packets []Packet) error
	// PauseResumeStream pauses or resumes the stream of sensor packets.
	PauseResumeStream() error
}

type roomba struct {
	port  io.ReadWriteCloser
	mutex sync.Mutex
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
func NewRoomba(port io.ReadWriteCloser) (Roomba, error) {
	if port == nil {
		return nil, errors.New("nil port supplied")
	}
	return &roomba{
		port: port,
	}, nil
}

func (r *roomba) Start() error {
	writeData := []byte{
		byte(OpCodeStart),
	}
	if err := r.writeAndRead(writeData, ProcessTimeStart, nil); err != nil {
		return err
	}
	return nil
}

func (r *roomba) SetBaudRate(baudRate BaudRate) error {
	_, ok := r.port.(serial.Port)
	if !ok {
		return errors.New("port not serial")
	}
	writeData := []byte{
		byte(OpCodeBaud),
		byte(baudRate),
	}
	if err := r.writeAndRead(writeData, ProcessTimeCommand, nil); err != nil {
		return err
	}
	return nil
}

func (r *roomba) Control() error {
	writeData := []byte{
		byte(OpCodeControl),
	}
	return r.writeAndRead(writeData, ProcessTimeModeChange, nil)
}

func (r *roomba) Safe() error {
	writeData := []byte{
		byte(OpCodeSafe),
	}
	return r.writeAndRead(writeData, ProcessTimeModeChange, nil)
}

func (r *roomba) Full() error {
	writeData := []byte{
		byte(OpCodeFull),
	}
	return r.writeAndRead(writeData, ProcessTimeModeChange, nil)
}

func (r *roomba) Power() error {
	writeData := []byte{
		byte(OpCodePower),
	}
	return r.writeAndRead(writeData, ProcessTimeModeChange, nil)
}

func (r *roomba) SetMode(mode Mode) error {
	switch mode {
	case ModeOff:
		return r.Power()
	case ModePassive:
		return r.Start()
	case ModeSafe:
		return r.Safe()
	case ModeFull:
		return r.Full()
	default:
		return errors.New("invalid mode")
	}
}

func (r *roomba) Spot() error {
	writeData := []byte{
		byte(OpCodeSpot),
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) Clean() error {
	writeData := []byte{
		byte(OpCodeClean),
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) Max() error {
	writeData := []byte{
		byte(OpCodeMax),
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) Drive(velocity int16, radius int16) error {
	if velocity < -500 || velocity > 500 {
		return fmt.Errorf("invalid velocity (%d)", velocity)
	}
	if radius < -2000 || radius > 2000 {
		return fmt.Errorf("invalid radius (%d)", radius)
	}
	vh, vl := int16ToBytes(velocity)
	rh, rl := int16ToBytes(radius)
	writeData := []byte{
		byte(OpCodeDrive),
		vh,
		vl,
		rh,
		rl,
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) Motors(mainBrush MainBrush, sideBrush SideBrush, vacuum Vacuum) error {
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
	writeData := []byte{
		byte(OpCodeMotors),
		bits,
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) LEDs(color uint8, intensity uint8, debris bool, spot bool, dock bool, checkRobot bool) error {
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
	writeData := []byte{
		byte(OpCodeLEDs),
		bits,
		color,
		intensity,
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
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
	writeData := make([]byte, 3+2*noteCount)
	writeData[0] = byte(OpCodeSong)
	writeData[1] = number
	writeData[2] = byte(len(notes))
	for i, note := range notes {
		writeData[3+i*2+0] = note.Number
		writeData[3+i*2+1] = note.Duration
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) Play(number uint8) error {
	if number > 4 {
		return fmt.Errorf("invalid song number (%d)", number)
	}
	writeData := []byte{
		byte(OpCodePlay),
		number,
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) Sensors(id int) (Packet, error) {
	packet, err := NewPacketWithID(id)
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
	writeData := []byte{
		byte(OpCodeSensors),
		byte(packet.ID()),
	}
	readData := make([]byte, packet.Size())
	// TODO
	if err := r.writeAndRead(writeData, ProcessTimeCommand, readData); err != nil {
		return err
	}
	return packet.Extract(readData, 0)
}

func (r *roomba) SeekDock() error {
	writeData := []byte{
		byte(OpCodeSeekDock),
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
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
	writeData := []byte{
		byte(OpCodeMotorsPWM),
		byte(mainBrushPWM),
		byte(sideBrushPWM),
		vacuumPWM,
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) DriveDirect(leftVelocity int16, rightVelocity int16) error {
	if leftVelocity < -500 || leftVelocity > 500 {
		return fmt.Errorf("invalid left velocity (%d)", leftVelocity)
	}
	if rightVelocity < -500 || rightVelocity > 500 {
		return fmt.Errorf("invalid right velocity (%d)", rightVelocity)
	}
	rvh, rvl := int16ToBytes(rightVelocity)
	lvh, lvl := int16ToBytes(leftVelocity)
	writeData := []byte{
		byte(OpCodeDriveDirect),
		rvh,
		rvl,
		lvh,
		lvl,
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) DrivePWM(leftWheelPWM int16, rightWheelPWM int16) error {
	if leftWheelPWM < -255 || leftWheelPWM > 255 {
		return fmt.Errorf("invalid left wheel PWM value (%d)", leftWheelPWM)
	}
	if rightWheelPWM < -255 || rightWheelPWM > 255 {
		return fmt.Errorf("invalid right wheel PWM value (%d)", rightWheelPWM)
	}
	rwh, rwl := int16ToBytes(rightWheelPWM)
	lwh, lwl := int16ToBytes(leftWheelPWM)
	writeData := []byte{
		byte(OpCodeDrivePWM),
		rwh,
		rwl,
		lwh,
		lwl,
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) Stream(ids ...int) (int, error) {
	if len(ids) > 255 {
		return 0, errors.New("too many sensor packets requested")
	}
	size := 0
	for _, id := range ids {
		packet, err := NewPacketWithID(id)
		if err != nil {
			return 0, err
		}
		size += packet.Size()
	}
	writeData := make([]byte, 2+len(ids))
	writeData[0] = byte(OpCodeStream)
	writeData[1] = byte(len(ids))
	for i, id := range ids {
		writeData[2+i] = byte(id)
	}
	if err := r.writeAndRead(writeData, ProcessTimeCommand, nil); err != nil {
		return 0, err
	}
	return size, nil
}

func (r *roomba) QueryList(ids ...int) ([]Packet, error) {
	packets := make([]Packet, 0)
	for _, id := range ids {
		packet, err := NewPacketWithID(id)
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)
	}
	if err := r.UpdateList(packets); err != nil {
		return nil, err
	}
	return packets, nil
}

func (r *roomba) UpdateList(packets []Packet) error {
	if len(packets) > 255 {
		return errors.New("too many sensor packets requested")
	}
	size := 0
	writeData := make([]byte, 2+len(packets))
	writeData[0] = byte(OpCodeQueryList)
	writeData[1] = byte(len(packets))
	for i, packet := range packets {
		writeData[2+i] = byte(packet.ID())
		size += packet.Size()
	}
	readData := make([]byte, size)
	if err := r.writeAndRead(writeData, ProcessTimeCommand, readData); err != nil {
		return err
	}
	offset := 0
	for _, packet := range packets {
		if err := packet.Extract(readData, offset); err != nil {
			return err
		}
		offset += packet.Size()
	}
	return nil
}

func (r *roomba) PauseResumeStream() error {
	writeData := []byte{
		byte(OpCodePauseResumeStream),
	}
	return r.writeAndRead(writeData, ProcessTimeCommand, nil)
}

func (r *roomba) Read(data []byte) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	n, err := r.port.Read(data)
	if err != nil {
		return 0, err
	}
	if n < len(data) {
		return 0, fmt.Errorf("incomplete read (%d)", n)
	}
	return n, nil
}

func (r *roomba) Write(data []byte) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	n, err := r.port.Write(data)
	if err != nil {
		return 0, err
	}
	if n < len(data) {
		return 0, fmt.Errorf("incomplete write (%d)", n)
	}
	return n, nil
}

func (r *roomba) Close() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.port.Close()
}

func (r *roomba) writeAndRead(writeData []byte, processTime time.Duration, readData []byte) error {
	if writeData != nil {
		if _, err := r.Write(writeData); err != nil {
			return err
		}
		time.Sleep(processTime)
	}
	if readData != nil {
		if _, err := r.Read(readData); err != nil {
			return err
		}
	}
	return nil
}
