// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package consensus

import "github.com/nebulasio/go-nebulas/core"

// Consensus interface of consensus algorithm.
type Consensus interface {
	Start()
	Stop()
	Event(e Event)
	TransiteByKey(nextStateKey string, data interface{})
	Transite(nextState State, data interface{})

	// AppendBlock add block to blockchain according to for choice algorithm.
	AppendBlock(block *core.Block) error
}
type Event interface{}

type State interface {
	Event(e Event) State
	Enter(data interface{})
	Leave(data interface{})
}

type States map[string]State

type BaseEvent struct {
	Data interface{}
}
