package main

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetTime(t *testing.T) {
	// Определяет время с сервера и хоста.
	ntpTime := getTime()
	hostTime := time.Now()

	// Разница между временами.
	different := hostTime.Sub(hostTime)

	// Если разница отрицательна, делает положительным
	if different < 0 {
		different = -different
	}

	// Если разница более в одну секунду, то тест провален.
	flag := false
	if different > time.Second {
		flag = true
	}

	// Сравнивание ожидаемого результата с полученным.
	assert.Equal(t, false, flag,
		fmt.Sprintf("Time from NTP: %+v\nTIme from host: %+v\nDifferent: %d", ntpTime, hostTime, different))
}