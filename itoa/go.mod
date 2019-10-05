module github.com/alexashley/cloud-native-fizzbuzz/itoa

go 1.12

require (
	github.com/alexashley/cloud-native-fizzbuzz/domain v0.0.0
	github.com/alexashley/cloud-native-fizzbuzz/server v0.0.0
)

replace github.com/alexashley/cloud-native-fizzbuzz/server => ../server

replace github.com/alexashley/cloud-native-fizzbuzz/domain => ../domain
