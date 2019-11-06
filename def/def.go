package def

import (
	"github.com/fancygo/fc_sdl"
)

const (
	P1 = 1
	P2 = 2

	D_U = 1
	D_R = 2
	D_D = 3
	D_L = 4

	LV = 100

	MAP_X = 26
	MAP_Y = 26

	WIN_W        = 670
	WIN_H        = 620
	MAP_W        = 520
	MAP_H        = 520
	U_REST       = 40
	L_REST       = 40
	R_REST       = 80
	D_REST       = 40
	TANK_W       = 40
	TANK_H       = 40
	TANK_SPEED   = 2
	TANK_LIFE    = 2
	TANK_LV_1    = 1
	TANK_LV_2    = 2
	TANK_LV_3    = 3
	TANK_LV_4    = 4
	BULLET_W     = 10
	BULLET_H     = 10
	BULLET_SPEED = 5
	BULLET_LV_1  = 1
	BULLET_LV_2  = 1

	COLLIDE_OFFSET = 4

	HOME_W = 40
	HOME_H = 40
	HOME_X = L_REST + MAP_W/2 - TERRAIN_W
	HOME_Y = U_REST + MAP_H - TANK_H

	TERRAIN_W = 20
	TERRAIN_H = 20

	TERRAIN_PASS_NO    = 0
	TERRAIN_PASS_YES   = 1
	TERRAIN_HIT_NO     = 0
	TERRAIN_HIT_NORMAL = 1
	TERRAIN_HIT_SUPER  = 2
	TERRAIN_SPEED_1    = 1
	TERRAIN_SPEED_2    = 2

	P1_BORN_X = L_REST + MAP_W/2 - TERRAIN_W*3 - TANK_W
	P1_BORN_Y = U_REST + MAP_H - TANK_H
	P2_BORN_X = L_REST + MAP_W/2 + TERRAIN_W*3
	P2_BORN_Y = U_REST + MAP_H - TANK_H

	E_NUM    = 3
	E_BORN_1 = L_REST
)

var E_POS [3]fc_sdl.Pos = [3]fc_sdl.Pos{
	{
		X: L_REST,
		Y: U_REST,
	},
	{
		X: WIN_W/2 - TANK_W/2,
		Y: U_REST,
	},
	{
		X: WIN_W - R_REST - TANK_W,
		Y: U_REST,
	},
}

const (
	ST_Null = iota
	ST_Player
	ST_Enermy
	ST_Home
	ST_Mud
	ST_Grass
	ST_Stone
	ST_Sea
	ST_Snow
	ST_Bullet
)
