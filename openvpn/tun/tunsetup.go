/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package tun

import "github.com/mysteriumnetwork/node/openvpn/config"

type linuxTunnelSetup struct {
}

func (lts *linuxTunnelSetup) Setup(config *config.GenericConfig) error {
	return nil
}

func (lts *linuxTunnelSetup) Teardown() {

}

type genericTunnelSetup struct {
}

func (gts *genericTunnelSetup) Setup(config *config.GenericConfig) error {
	return nil
}

func (gts *genericTunnelSetup) Teardown() {

}
