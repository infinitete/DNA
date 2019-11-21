/*
 * Copyright (C) 2018 The DNA Authors
 * This file is part of The DNA library.
 *
 * The DNA is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The DNA is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The DNA.  If not, see <http://www.gnu.org/licenses/>.
 */

package message

import (
	"git.fe-cred.com/idfor/idfor/common"
	"git.fe-cred.com/idfor/idfor/core/types"
)

const (
	TOPIC_SAVE_BLOCK_COMPLETE       = "svblkcmp"
	TOPIC_NEW_INVENTORY             = "newinv"
	TOPIC_NODE_DISCONNECT           = "noddis"
	TOPIC_NODE_CONSENSUS_DISCONNECT = "nodcnsdis"
	TOPIC_SMART_CODE_EVENT          = "scevt"
)

type SaveBlockCompleteMsg struct {
	Block *types.Block
}

type NewInventoryMsg struct {
	Inventory *common.Inventory
}

type SmartCodeEventMsg struct {
	Event *types.SmartCodeEvent
}

type BlockConsensusComplete struct {
	Block *types.Block
}
