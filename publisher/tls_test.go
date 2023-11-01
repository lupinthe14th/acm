package main

import (
	"testing"
)

const fakeCA = "emqxsl-ca.crt"

func TestNewTLSConfig(t *testing.T) {
	tlsConfig, err := newTLSConfig("")
	if err != nil {
		t.Errorf("Could not create TLS config: %v", err)
	}

	// RootCAsが設定されているか確認
	if tlsConfig.RootCAs == nil {
		t.Fatal("Expected RootCAs to be set")
	}

	// InsecureSkipVerifyがfalseであるか確認
	if tlsConfig.InsecureSkipVerify {
		t.Fatal("Expected InsecureSkipVerify to be false")
	}
}

func TestNewTLSConfigAddFile(t *testing.T) {
	tlsConfig, err := newTLSConfig(fakeCA)
	if err != nil {
		t.Errorf("Could not create TLS config: %v", err)
	}

	// RootCAsが設定されているか確認
	if tlsConfig.RootCAs == nil {
		t.Fatal("Expected RootCAs to be set")
	}

	// 証明書が正しくcertpoolに追加されているか確認
	if len(tlsConfig.RootCAs.Subjects()) == 0 {
		t.Fatal("No certificates found in RootCAs")
	}

	// InsecureSkipVerifyがfalseであるか確認
	if tlsConfig.InsecureSkipVerify {
		t.Fatal("Expected InsecureSkipVerify to be false")
	}
}
