package cachekey

const cacheKeyPrefix string = "cache:landisland:key:"

func Account(userId string) string {
	return cacheKeyPrefix + "account" + userId
}
