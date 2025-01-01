package config

const Content = ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_'abcdefjhijklmnopqrstuvwxyz{|}~`
var Standard =  make(map[rune][]string, len(Content))
var Shadow =  make(map[rune][]string, len(Content))
var Thinkertoy =  make(map[rune][]string, len(Content))
const CharacterHeight = 8