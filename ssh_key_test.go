// This file is part of gobizfly
//
// Copyright (C) 2020  BizFly Cloud
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package gobizfly

import (
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestSSHKeyList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.CloudServerURL(sshKeyBasePath), func(writer http.ResponseWriter, request *http.Request) {
		require.Equal(t, http.MethodGet, request.Method)
		resp := `
[
    {
        "keypair": {
            "name": "sapd1",
            "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDgXX0Kdd3XKojgj7maVd3PsPApzh9n2lT2CtgcJs8jw9i3mit5SZu02QFS772Pa9VdGeSjbqxtADLRpnuigW5ii0dHBQTgWqx593Cs7QKRhyRPb88u0TFCZynRwfMRnb6qngiKoWp5TtaHuIY+7kS8SyqNVIwoCYlr9a4ePX8rwydf9crhJocgKb2LgQkdW3TBE5QAvxbruYlj201jjXFeE5BtE4QER0QyY5MqW8MAgG98N3w95pKIffhHZ4TO4A3zgpWbNn1ROproZgV+9COzZ7WYuvPWqWdLAntd9b1/lLnDrDHXa/lrefJXJVamhz4i1cfIZ/p+aFWG0a7DpL5b saphi@saphi-kma\n",
            "fingerprint": "28:56:9e:4b:bb:a0:91:71:42:37:40:a2:d0:66:24:17"
        }
    },
    {
        "keypair": {
            "name": "ssh-key-1601308682626",
            "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC5Dl5SGjkhI/RhBpqxzPyL1sEfkte1OWUMgF5xwRJWrhvGDZ8LTUBUxE5cBBxRSccR+tdU5R66K5JxdcmbyAFDXEXqUJbKZj8mF/wCYa3JzTLJaCS/pys4Mx6+59kpAREROr5hvnm9HLtvVS7MlBxFh/tIRCWMsVhRdVpK22TFhSNGVxR3Xc2OsIhq119HCIpApGae5tlvHq4Kn+EZs4DCaXY0dUpNdaJbJtN0TiQFJl3/NxV9s6VRiKdy643EymvFC4TaurSVYW2H9Tr3PYkxQrXcCyo6ZMQrCFeijFyYXDf4gfQu9iWKk71ZZYQPXz1ThS31JQkdB/T96h4C6PLz Generated-by-Nova",
            "fingerprint": "75:65:c3:da:7e:97:72:90:26:44:4c:a1:f9:0f:01:51"
        }
    }
]
`
		_, _ = fmt.Fprint(writer, resp)
	})
	sshkeys, err :=client.SSHKey.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, sshkeys, 2)
	assert.Equal(t, "sapd1", sshkeys[0].SSHKeyPair.Name)
	assert.Equal(t, "75:65:c3:da:7e:97:72:90:26:44:4c:a1:f9:0f:01:51", sshkeys[1].SSHKeyPair.FingerPrint)
}

func TestSSHKeyDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.CloudServerURL(sshKeyBasePath + "/sapd1"), func(writer http.ResponseWriter, request *http.Request) {
		require.Equal(t, http.MethodDelete, request.Method)
		resp := `
{
	"message": "Delete successful"
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	_, err :=client.SSHKey.Delete(ctx, "sapd1")
	require.NoError(t, err)
}

func TestSSHKeyCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.CloudServerURL(sshKeyBasePath), func(writer http.ResponseWriter, request *http.Request) {
		require.Equal(t, http.MethodPost, request.Method)
		resp := `
{
    "name": "ssh-key-1601308814384",
    "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDgXX0Kdd3XKojgj7maVd3PsPApzh9n2lT2CtgcJs8jw9i3mit5SZu02QFS772Pa9VdGeSjbqxtADLRpnuigW5ii0dHBQTgWqx593Cs7QKRhyRPb88u0TFCZynRwfMRnb6qngiKoWp5TtaHuIY+7kS8SyqNVIwoCYlr9a4ePX8rwydf9crhJocgKb2LgQkdW3TBE5QAvxbruYlj201jjXFeE5BtE4QER0QyY5MqW8MAgG98N3w95pKIffhHZ4TO4A3zgpWbNn1ROproZgV+9COzZ7WYuvPWqWdLAntd9b1/lLnDrDHXa/lrefJXJVamhz4i1cfIZ/p+aFWG0a7DpL5b saphi@saphi-kma\n",
    "fingerprint": "28:56:9e:4b:bb:a0:91:71:42:37:40:a2:d0:66:24:17",
    "user_id": "55d38aecb1034c06b99c1c87fb6f0740"
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	sshkey, err :=client.SSHKey.Create(ctx, &SSHKeyCreateRequest{
		Name: "ssh-key-1601308814384",
		PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDgXX0Kdd3XKojgj7maVd3PsPApzh9n2lT2CtgcJs8jw9i3mit5SZu02QFS772Pa9VdGeSjbqxtADLRpnuigW5ii0dHBQTgWqx593Cs7QKRhyRPb88u0TFCZynRwfMRnb6qngiKoWp5TtaHuIY+7kS8SyqNVIwoCYlr9a4ePX8rwydf9crhJocgKb2LgQkdW3TBE5QAvxbruYlj201jjXFeE5BtE4QER0QyY5MqW8MAgG98N3w95pKIffhHZ4TO4A3zgpWbNn1ROproZgV+9COzZ7WYuvPWqWdLAntd9b1/lLnDrDHXa/lrefJXJVamhz4i1cfIZ/p+aFWG0a7DpL5b saphi@saphi-kma\n",
	})
	require.NoError(t, err)
	assert.Equal(t, "ssh-key-1601308814384", sshkey.Name)
	assert.Equal(t, "55d38aecb1034c06b99c1c87fb6f0740", sshkey.UserID)
}