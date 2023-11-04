package main

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name                 string
		serverURL            string
		caFile               string
		clientID             string
		username             string
		password             string
		topic                string
		qos                  string
		keepAlive            string
		connectRetryDelay    string
		delayBetweenMessages string
		printMessages        string
		debug                string
		wantConfig           config
		isErr                bool
	}{
		{
			name:                 "standerd case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig: config{
				serverURL: &url.URL{
					Scheme: "http",
					Host:   "localhost:1883"},
				caFile:               "ca.pem",
				clientID:             "publisher00001",
				username:             "user",
				password:             "pass",
				topic:                "/example/#",
				qos:                  byte(0),
				keepAlive:            30,
				connectRetryDelay:    time.Duration(30) * time.Millisecond,
				delayBetweenMessages: time.Duration(15) * time.Millisecond,
				printMessage:         true,
				debug:                false,
			},
			isErr: false,
		},
		{
			name:                 "serverURL must not be blank case",
			serverURL:            "",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "must be a valid URL case",
			serverURL:            ":",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "caFile must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "clientID must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "topic must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "qos must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "keepAlive must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "connectRetryDelay  must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "delayBetweenMessages must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "",
			printMessages:        "true",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "printMessages must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "",
			debug:                "false",
			wantConfig:           config{},
			isErr:                true,
		},
		{
			name:                 "debug must not be blank case",
			serverURL:            "http://localhost:1883",
			caFile:               "ca.pem",
			clientID:             "publisher00001",
			username:             "user",
			password:             "pass",
			topic:                "/example/#",
			qos:                  "0",
			keepAlive:            "30",
			connectRetryDelay:    "30",
			delayBetweenMessages: "15",
			printMessages:        "true",
			debug:                "",
			wantConfig:           config{},
			isErr:                true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("acm_serverURL", tt.serverURL)
			t.Setenv("acm_caFile", tt.caFile)
			t.Setenv("acm_clientID", tt.clientID)
			t.Setenv("acm_username", tt.username)
			t.Setenv("acm_password", tt.password)
			t.Setenv("acm_topic", tt.topic)
			t.Setenv("acm_qos", tt.qos)
			t.Setenv("acm_keepAlive", tt.keepAlive)
			t.Setenv("acm_connectRetryDelay", tt.connectRetryDelay)
			t.Setenv("acm_delayBetweenMessages", tt.delayBetweenMessages)
			t.Setenv("acm_printMessages", tt.printMessages)
			t.Setenv("acm_debug", tt.debug)
			got, err := getConfig()
			if !reflect.DeepEqual(got, tt.wantConfig) {
				t.Fatalf("unexpected value: got: %v, want: %v", got, tt.wantConfig)
			}
			if tt.isErr && err == nil {
				t.Fatalf("unexpected error: got: %v, want: %t", err, tt.isErr)
			}
		})
	}
}

func TestStringFromEnv(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		wantValue string
		isErr     bool
	}{
		{
			name:      "get value",
			key:       "smallpox",
			value:     "SMALLPOX",
			wantValue: "SMALLPOX",
			isErr:     false,
		},
		{
			name:      "must not be blank",
			key:       "smallpox",
			value:     "",
			wantValue: "",
			isErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.key, tt.value)
			got, err := stringFromEnv(tt.key)
			if got != tt.wantValue {
				t.Fatalf("unexpected value: got: %s, want: %s", got, tt.wantValue)
			}
			if tt.isErr && err == nil {
				t.Fatalf("unexpected error: got: %v, want: %t", err, tt.isErr)
			}
		})
	}
}

func TestIntFromEnv(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		wantValue int
		isErr     bool
	}{
		{
			name:      "get value: standerd case",
			key:       "smallpox",
			value:     "1",
			wantValue: 1,
			isErr:     false,
		},
		{
			name:      "get value: corner case",
			key:       "smallpox",
			value:     "0",
			wantValue: 0,
			isErr:     false,
		},
		{
			name:      "must not be blank",
			key:       "smallpox",
			value:     "",
			wantValue: 0,
			isErr:     true,
		},
		{
			name:      "must not be integer",
			key:       "smallpox",
			value:     "a",
			wantValue: 0,
			isErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.key, tt.value)
			got, err := intFromEnv(tt.key)
			if got != tt.wantValue {
				t.Fatalf("unexpected value: got: %d, want: %d", got, tt.wantValue)
			}
			if tt.isErr && err == nil {
				t.Fatalf("unexpected error: got: %v, want: %t", err, tt.isErr)
			}
		})
	}
}

func TestMillSeconfFromEnv(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		wantValue time.Duration
		isErr     bool
	}{
		{
			name:      "get value: standerd case",
			key:       "smallpox",
			value:     "1",
			wantValue: time.Duration(1) * time.Millisecond,
			isErr:     false,
		},
		{
			name:      "get value: corner case",
			key:       "smallpox",
			value:     "0",
			wantValue: 0,
			isErr:     false,
		},
		{
			name:      "must not be blank",
			key:       "smallpox",
			value:     "",
			wantValue: 0,
			isErr:     true,
		},
		{
			name:      "must not be integer",
			key:       "smallpox",
			value:     "a",
			wantValue: 0,
			isErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.key, tt.value)
			got, err := milliSecondsFromEnv(tt.key)
			if got != tt.wantValue {
				t.Fatalf("unexpected value: got: %d, want: %d", got, tt.wantValue)
			}
			if tt.isErr && err == nil {
				t.Fatalf("unexpected error: got: %v, want: %t", err, tt.isErr)
			}
		})
	}
}
func TestBoolFromEnv(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		wantValue bool
		isErr     bool
	}{
		{
			name:      "get value: true case: TRUE",
			key:       "smallpox",
			value:     "TRUE",
			wantValue: true,
			isErr:     false,
		},
		{
			name:      "get value: true case: true",
			key:       "smallpox",
			value:     "true",
			wantValue: true,
			isErr:     false,
		},
		{
			name:      "get value: true case: T",
			key:       "smallpox",
			value:     "T",
			wantValue: true,
			isErr:     false,
		},
		{
			name:      "get value: true case: t",
			key:       "smallpox",
			value:     "t",
			wantValue: true,
			isErr:     false,
		},
		{
			name:      "get value: true case: 1",
			key:       "smallpox",
			value:     "1",
			wantValue: true,
			isErr:     false,
		},
		{
			name:      "get value: false case: FALSE",
			key:       "smallpox",
			value:     "FALSE",
			wantValue: false,
			isErr:     false,
		},
		{
			name:      "get value: false case: false",
			key:       "smallpox",
			value:     "false",
			wantValue: false,
			isErr:     false,
		},
		{
			name:      "get value: false case: F",
			key:       "smallpox",
			value:     "F",
			wantValue: false,
			isErr:     false,
		},
		{
			name:      "get value: false case: f",
			key:       "smallpox",
			value:     "f",
			wantValue: false,
			isErr:     false,
		},
		{
			name:      "get value: false case: 0",
			key:       "smallpox",
			value:     "0",
			wantValue: false,
			isErr:     false,
		},
		{
			name:      "must not be blank",
			key:       "smallpox",
			value:     "",
			wantValue: false,
			isErr:     true,
		},
		{
			name:      "be a valid boolean option",
			key:       "smallpox",
			value:     "2",
			wantValue: false,
			isErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.key, tt.value)
			got, err := booleanFromEnv(tt.key)
			if got != tt.wantValue {
				t.Fatalf("unexpected value: got: %v, want: %v", got, tt.wantValue)
			}
			if tt.isErr && err == nil {
				t.Fatalf("unexpected error: got: %v, want: %t", err, tt.isErr)
			}
		})
	}
}
