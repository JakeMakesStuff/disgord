package disgord

// Code generated - This file has been automatically generated by generate/events/main.go - DO NOT EDIT.
// Warning: This file is overwritten at "go generate", instead adapt events.go and event/cache.go and run go generate

import (
	"github.com/andersfylling/disgord/internal/event"
	"github.com/andersfylling/disgord/json"
)

// Cache interface for event handling and REST methods
// commented out fields are simply not supported yet. PR's are welcome
//
// Note that on events you are expected to return a unmarshalled object. For delete methods
// you should return nil, and a nil error if the objected to be deleted was not found (nop!).
// Note that the error might change to a "CacheMiss" or something similar such that we can
//  get more metrics!
type Cache interface {
	CacheUpdater
	CacheGetter
}

type CacheGetter interface {
	// REST API

	// GetGuildAuditLogs(guildID Snowflake) *guildAuditLogsBuilder // TODO
	GetMessages(channelID Snowflake, params *GetMessagesParams) ([]*Message, error)
	GetMessage(channelID, messageID Snowflake) (ret *Message, err error)
	//GetUsersWhoReacted(channelID, messageID Snowflake, emoji interface{}, params URLQueryStringer) (reactors []*User, err error)
	//GetPinnedMessages(channelID Snowflake) (ret []*Message, err error)
	GetChannel(id Snowflake) (ret *Channel, err error)
	//GetChannelInvites(id Snowflake) (ret []*Invite, err error)
	GetGuildEmoji(guildID, emojiID Snowflake) (*Emoji, error)
	GetGuildEmojis(id Snowflake) ([]*Emoji, error)
	GetGuild(id Snowflake) (*Guild, error)
	GetGuildChannels(id Snowflake) ([]*Channel, error)
	GetMember(guildID, userID Snowflake) (*Member, error)
	GetMembers(guildID Snowflake, params *GetMembersParams) ([]*Member, error)
	//GetGuildBans(id Snowflake) ([]*Ban, error)
	//GetGuildBan(guildID, userID Snowflake) (*Ban, error)
	GetGuildRoles(guildID Snowflake) ([]*Role, error)
	//GetMemberPermissions(guildID, userID Snowflake) (permissions PermissionBit, err error)
	//GetGuildVoiceRegions(id Snowflake) ([]*VoiceRegion, error)
	//GetGuildInvites(id Snowflake) ([]*Invite, error)
	//GetGuildIntegrations(id Snowflake) ([]*Integration, error)
	//GetGuildEmbed(guildID Snowflake) (*GuildEmbed, error)
	//GetGuildVanityURL(guildID Snowflake) (*PartialInvite, error)
	//GetInvite(inviteCode string, params URLQueryStringer) (*Invite, error)
	GetCurrentUser() (*User, error)
	GetUser(id Snowflake) (*User, error)
	GetCurrentUserGuilds(params *GetCurrentUserGuildsParams) (ret []*Guild, err error)
	//GetUserDMs() (ret []*Channel, err error)
	//GetUserConnections() (ret []*UserConnection, err error)
	//GetVoiceRegions() ([]*VoiceRegion, error)
	//GetChannelWebhooks(channelID Snowflake) (ret []*Webhook, err error)
	//GetGuildWebhooks(guildID Snowflake) (ret []*Webhook, err error)
	//GetWebhook(id Snowflake) (ret *Webhook, err error)
}

type CacheUpdater interface {
	// Gateway events
	ChannelCreate(data []byte) (*ChannelCreate, error)
	ChannelDelete(data []byte) (*ChannelDelete, error)
	ChannelPinsUpdate(data []byte) (*ChannelPinsUpdate, error)
	ChannelUpdate(data []byte) (*ChannelUpdate, error)
	GuildBanAdd(data []byte) (*GuildBanAdd, error)
	GuildBanRemove(data []byte) (*GuildBanRemove, error)
	GuildCreate(data []byte) (*GuildCreate, error)
	GuildDelete(data []byte) (*GuildDelete, error)
	GuildEmojisUpdate(data []byte) (*GuildEmojisUpdate, error)
	GuildIntegrationsUpdate(data []byte) (*GuildIntegrationsUpdate, error)
	GuildMemberAdd(data []byte) (*GuildMemberAdd, error)
	GuildMemberRemove(data []byte) (*GuildMemberRemove, error)
	GuildMemberUpdate(data []byte) (*GuildMemberUpdate, error)
	GuildMembersChunk(data []byte) (*GuildMembersChunk, error)
	GuildRoleCreate(data []byte) (*GuildRoleCreate, error)
	GuildRoleDelete(data []byte) (*GuildRoleDelete, error)
	GuildRoleUpdate(data []byte) (*GuildRoleUpdate, error)
	GuildUpdate(data []byte) (*GuildUpdate, error)
	InviteCreate(data []byte) (*InviteCreate, error)
	InviteDelete(data []byte) (*InviteDelete, error)
	MessageCreate(data []byte) (*MessageCreate, error)
	MessageDelete(data []byte) (*MessageDelete, error)
	MessageDeleteBulk(data []byte) (*MessageDeleteBulk, error)
	MessageReactionAdd(data []byte) (*MessageReactionAdd, error)
	MessageReactionRemove(data []byte) (*MessageReactionRemove, error)
	MessageReactionRemoveAll(data []byte) (*MessageReactionRemoveAll, error)
	MessageReactionRemoveEmoji(data []byte) (*MessageReactionRemoveEmoji, error)
	MessageUpdate(data []byte) (*MessageUpdate, error)
	PresenceUpdate(data []byte) (*PresenceUpdate, error)
	Ready(data []byte) (*Ready, error)
	Resumed(data []byte) (*Resumed, error)
	TypingStart(data []byte) (*TypingStart, error)
	UserUpdate(data []byte) (*UserUpdate, error)
	VoiceServerUpdate(data []byte) (*VoiceServerUpdate, error)
	VoiceStateUpdate(data []byte) (*VoiceStateUpdate, error)
	WebhooksUpdate(data []byte) (*WebhooksUpdate, error)
}

func cacheDispatcher(c Cache, eventName string, data []byte) (evt EventType, err error) {
	switch eventName {
	case event.ChannelCreate:
		evt, err = c.ChannelCreate(data)
	case event.ChannelDelete:
		evt, err = c.ChannelDelete(data)
	case event.ChannelPinsUpdate:
		evt, err = c.ChannelPinsUpdate(data)
	case event.ChannelUpdate:
		evt, err = c.ChannelUpdate(data)
	case event.GuildBanAdd:
		evt, err = c.GuildBanAdd(data)
	case event.GuildBanRemove:
		evt, err = c.GuildBanRemove(data)
	case event.GuildCreate:
		evt, err = c.GuildCreate(data)
	case event.GuildDelete:
		evt, err = c.GuildDelete(data)
	case event.GuildEmojisUpdate:
		evt, err = c.GuildEmojisUpdate(data)
	case event.GuildIntegrationsUpdate:
		evt, err = c.GuildIntegrationsUpdate(data)
	case event.GuildMemberAdd:
		evt, err = c.GuildMemberAdd(data)
	case event.GuildMemberRemove:
		evt, err = c.GuildMemberRemove(data)
	case event.GuildMemberUpdate:
		evt, err = c.GuildMemberUpdate(data)
	case event.GuildMembersChunk:
		evt, err = c.GuildMembersChunk(data)
	case event.GuildRoleCreate:
		evt, err = c.GuildRoleCreate(data)
	case event.GuildRoleDelete:
		evt, err = c.GuildRoleDelete(data)
	case event.GuildRoleUpdate:
		evt, err = c.GuildRoleUpdate(data)
	case event.GuildUpdate:
		evt, err = c.GuildUpdate(data)
	case event.InviteCreate:
		evt, err = c.InviteCreate(data)
	case event.InviteDelete:
		evt, err = c.InviteDelete(data)
	case event.MessageCreate:
		evt, err = c.MessageCreate(data)
	case event.MessageDelete:
		evt, err = c.MessageDelete(data)
	case event.MessageDeleteBulk:
		evt, err = c.MessageDeleteBulk(data)
	case event.MessageReactionAdd:
		evt, err = c.MessageReactionAdd(data)
	case event.MessageReactionRemove:
		evt, err = c.MessageReactionRemove(data)
	case event.MessageReactionRemoveAll:
		evt, err = c.MessageReactionRemoveAll(data)
	case event.MessageReactionRemoveEmoji:
		evt, err = c.MessageReactionRemoveEmoji(data)
	case event.MessageUpdate:
		evt, err = c.MessageUpdate(data)
	case event.PresenceUpdate:
		evt, err = c.PresenceUpdate(data)
	case event.Ready:
		evt, err = c.Ready(data)
	case event.Resumed:
		evt, err = c.Resumed(data)
	case event.TypingStart:
		evt, err = c.TypingStart(data)
	case event.UserUpdate:
		evt, err = c.UserUpdate(data)
	case event.VoiceServerUpdate:
		evt, err = c.VoiceServerUpdate(data)
	case event.VoiceStateUpdate:
		evt, err = c.VoiceStateUpdate(data)
	case event.WebhooksUpdate:
		evt, err = c.WebhooksUpdate(data)
	default:
		evt = nil
		err = nil
	}

	return evt, err
}

// nop cache
type CacheNop struct{}

var _ CacheUpdater = (*CacheNop)(nil)
var _ CacheGetter = (*CacheNop)(nil)
var _ Cache = (*CacheNop)(nil)

func (c *CacheNop) Patch(v interface{}) {
	if v == nil {
		return
	}
	if deficient, ok := v.(internalUpdater); ok {
		deficient.updateInternals()
	}
}
func (c *CacheNop) ChannelCreate(data []byte) (evt *ChannelCreate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) ChannelDelete(data []byte) (evt *ChannelDelete, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) ChannelPinsUpdate(data []byte) (evt *ChannelPinsUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) ChannelUpdate(data []byte) (evt *ChannelUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildBanAdd(data []byte) (evt *GuildBanAdd, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildBanRemove(data []byte) (evt *GuildBanRemove, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildCreate(data []byte) (evt *GuildCreate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildDelete(data []byte) (evt *GuildDelete, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildEmojisUpdate(data []byte) (evt *GuildEmojisUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildIntegrationsUpdate(data []byte) (evt *GuildIntegrationsUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildMemberAdd(data []byte) (evt *GuildMemberAdd, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildMemberRemove(data []byte) (evt *GuildMemberRemove, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildMemberUpdate(data []byte) (evt *GuildMemberUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildMembersChunk(data []byte) (evt *GuildMembersChunk, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildRoleCreate(data []byte) (evt *GuildRoleCreate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildRoleDelete(data []byte) (evt *GuildRoleDelete, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildRoleUpdate(data []byte) (evt *GuildRoleUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) GuildUpdate(data []byte) (evt *GuildUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) InviteCreate(data []byte) (evt *InviteCreate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) InviteDelete(data []byte) (evt *InviteDelete, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) MessageCreate(data []byte) (evt *MessageCreate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) MessageDelete(data []byte) (evt *MessageDelete, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) MessageDeleteBulk(data []byte) (evt *MessageDeleteBulk, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) MessageReactionAdd(data []byte) (evt *MessageReactionAdd, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) MessageReactionRemove(data []byte) (evt *MessageReactionRemove, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) MessageReactionRemoveAll(data []byte) (evt *MessageReactionRemoveAll, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) MessageReactionRemoveEmoji(data []byte) (evt *MessageReactionRemoveEmoji, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) MessageUpdate(data []byte) (evt *MessageUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) PresenceUpdate(data []byte) (evt *PresenceUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) Ready(data []byte) (evt *Ready, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) Resumed(data []byte) (evt *Resumed, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) TypingStart(data []byte) (evt *TypingStart, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) UserUpdate(data []byte) (evt *UserUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) VoiceServerUpdate(data []byte) (evt *VoiceServerUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) VoiceStateUpdate(data []byte) (evt *VoiceStateUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}
func (c *CacheNop) WebhooksUpdate(data []byte) (evt *WebhooksUpdate, err error) {
	if err = json.Unmarshal(data, &evt); err != nil {
		return nil, err
	}
	c.Patch(evt)
	return evt, nil
}

// REST lookup
func (c *CacheNop) GetMessage(channelID, messageID Snowflake) (*Message, error) { return nil, nil }
func (c *CacheNop) GetChannel(id Snowflake) (*Channel, error)                   { return nil, nil }
func (c *CacheNop) GetGuildEmoji(guildID, emojiID Snowflake) (*Emoji, error)    { return nil, nil }
func (c *CacheNop) GetGuildEmojis(id Snowflake) ([]*Emoji, error)               { return nil, nil }
func (c *CacheNop) GetGuild(id Snowflake) (*Guild, error)                       { return nil, nil }
func (c *CacheNop) GetGuildChannels(id Snowflake) ([]*Channel, error)           { return nil, nil }
func (c *CacheNop) GetMember(guildID, userID Snowflake) (*Member, error)        { return nil, nil }
func (c *CacheNop) GetGuildRoles(guildID Snowflake) ([]*Role, error)            { return nil, nil }
func (c *CacheNop) GetCurrentUser() (*User, error)                              { return nil, nil }
func (c *CacheNop) GetUser(id Snowflake) (*User, error)                         { return nil, nil }
func (c *CacheNop) GetCurrentUserGuilds(p *GetCurrentUserGuildsParams) ([]*Guild, error) {
	return nil, nil
}
func (c *CacheNop) GetMessages(channel Snowflake, p *GetMessagesParams) ([]*Message, error) {
	return nil, nil
}
func (c *CacheNop) GetMembers(guildID Snowflake, p *GetMembersParams) ([]*Member, error) {
	return nil, nil
}
