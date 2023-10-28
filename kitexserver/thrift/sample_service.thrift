namespace go api

include "dragon.thrift"

service Sample {
    dragon.DragonSayResponse dragonSay(1: dragon.DragonSayRequest req)
}