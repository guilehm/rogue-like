package settings

import "time"

const MoveStep = 1
const MoveRange = 8
const PercentageToAttackBack = 20
const RespawnCheckTime = 1 * time.Second
const FollowCheckTime = 250 * time.Millisecond
const IncreasePlayersHealthValue = 10
const IncreasePlayersHealthCheckTime = 5 * time.Second
const ViewAreaOffsetX = 8 * 9
const ViewAreaOffsetY = 8 * 6
const ProjectileMoveTime = 20 * time.Millisecond
const NextLevelXpIncreaseRate = 1.2
const BaseNextLevelXP = 30
const CooldownForRespawnedEnemies = 3 * time.Second
