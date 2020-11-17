local Constants = require("packet/Constants")
local Helpers = require("packet/Helpers")
local PacketTypes = require("packet/PacketTypes")
local ServerClientCommands = require("packet/ServerClientCommands")

-- lookup tables
local commands_by_id = {}
local commands_by_key = {}

for _, cmd in ipairs(ServerClientCommands) do
	commands_by_id[cmd.id] = cmd
	commands_by_key[cmd.key] = cmd
end

local packet_types_by_id = {}
for k, id in pairs(PacketTypes) do
	packet_types_by_id[id] = k
end


local function parse(buf)
	assert(#buf > 10, "packet too small")

	for i=1,4 do
		assert(string.byte(Constants.protocol_id, i) == string.byte(buf, i))
	end

	-- 4F457403 0001 00 03 FFDC 0001 00F4
	-- 4F457403 0001 00 00 00 FFDC
	-- 4F457403 0001 00 03 FFDD 0002

	local def = {}

	def.peer_id = Helpers.bytes_to_int( string.byte(buf, 5), string.byte(buf, 6) )
	def.channel = string.byte(buf, 7)

	def.type = packet_types_by_id[string.byte(buf, 8)]

	if def.type == "reliable" then
		def.sequence_nr = Helpers.bytes_to_int( string.byte(buf, 9), string.byte(buf, 10) )
		def.command_id = Helpers.bytes_to_int( string.byte(buf, 11), string.byte(buf, 12) )
		local cmd = commands_by_id[def.command_id]
		def.command = cmd.key
		def.payload = cmd.parse(buf:sub(13))

	elseif def.type == "original" then
		error("not implemented")

	elseif def.type == "control" then
		if string.byte(buf, 9) == 0 then
			-- ack
			def.ack = true
			def.sequence_nr = Helpers.bytes_to_int( string.byte(buf, 10), string.byte(buf, 11) )
		else
			error("not implemented")
		end

	end


	return def
end

local function create()
	-- TODO
end

-- exports
return {
	create = create,
	parse = parse
}
