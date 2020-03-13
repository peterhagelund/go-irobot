package irobot

import (
	"errors"
	"net"
	"time"
)

// OpCode is the OIC op-code type.
type OpCode byte

const (
	// OpCodeStart stats the OIC.
	OpCodeStart OpCode = iota + 128
	// OpCodeBaud changes the baud rate of the OIC.
	OpCodeBaud
	_
	// OpCodeSafe puts the OIC in safe mode.
	OpCodeSafe
	// OpCodeFull puts the OIC in full mode.
	OpCodeFull
	// OpCodePower powers off the OIC.
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
)

// BaudRate is the OIC baud rate type.
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

// Roomba defines the required behavior of an iRobot Roomba vacuum cleaner OIC.
type Roomba interface {
	// Start starts the OIC.
	Start() error
	// SetBaud sets the baud rate of the OIC.
	SetBaud(baudRate BaudRate) error
	// Safe puts the OIC in safe mode.
	Safe() error
	// Full puts the OIC in full mode.
	Full() error
	// Power powers off the OIC.
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

	Sensors(id int) (Packet, error)
}

type roomba struct {
	conn      net.Conn
	lastWrite time.Time
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
	data := [...]byte{byte(OpCodeStart)}
	return r.write(data[:])
}

func (r *roomba) SetBaud(baudRate BaudRate) error {
	data := [...]byte{byte(OpCodeBaud), byte(baudRate)}
	return r.write(data[:])
}

func (r *roomba) Safe() error {
	data := [...]byte{byte(OpCodeSafe)}
	return r.write(data[:])
}

func (r *roomba) Full() error {
	data := [...]byte{byte(OpCodeFull)}
	return r.write(data[:])
}

func (r *roomba) Power() error {
	data := [...]byte{byte(OpCodePower)}
	return r.write(data[:])
}

func (r *roomba) Spot() error {
	data := [...]byte{byte(OpCodeSpot)}
	return r.write(data[:])
}

func (r *roomba) Clean() error {
	data := [...]byte{byte(OpCodeClean)}
	return r.write(data[:])
}

func (r *roomba) Max() error {
	data := [...]byte{byte(OpCodeMax)}
	return r.write(data[:])
}

func (r *roomba) Drive(velocity int16, radius int16) error {
	return nil
}

func (r *roomba) Motors(mainBrush MainBrush, sideBrush SideBrush, vacuum Vacuum) error {
	return nil
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

func (r *roomba) write(data []byte) error {
	for {
		if r.lastWrite.Add(time.Duration(50 * time.Millisecond)).Before(time.Now()) {
			break
		}
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	n, err := r.conn.Write(data)
	if err != nil {
		return err
	}
	if n < len(data) {
		return errors.New("incomplete write")
	}
	return nil
}
