package model

import "hash/crc64"

const AlphabetBase63 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

var HashTable = crc64.MakeTable(crc64.ECMA)
