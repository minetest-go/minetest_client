
local Constants = require("packet/Constants")
local Helpers = require("packet/Helpers")
local PacketTypes = require("packet/PacketTypes")
local ClientServerCommands = require("packet/ClientServerCommands")

-- lookup tables
local packet_types_by_id = {}
for k, id in pairs(PacketTypes) do
	packet_types_by_id[id] = k
end

local commands_by_id = {}
local commands_by_key = {}

for _, cmd in ipairs(ClientServerCommands) do
	commands_by_id[cmd.id] = cmd
	commands_by_key[cmd.key] = cmd
end

local function create(def)
	-- assert(commands[def.command], "command not available: " .. def.command)

	local packet = 	Constants.protocol_id ..
		Helpers.int_to_bytes(def.peer_id) .. -- peer_id
		string.char(def.channel) .. -- channel
		string.char(PacketTypes[def.type]) -- type


	if def.type == "reliable" then
		packet = packet .. Helpers.int_to_bytes(def.sequence_nr) .. -- seq nr
			string.char(PacketTypes[def.subtype]) .. -- subtype
			def.payload

	elseif def.type == "original" then
		-- command + payload
		local cmd = commands_by_key[def.command]

		packet = packet ..
			Helpers.int_to_bytes(cmd.id) ..
			cmd.create(def.payload)

	elseif def.type == "control" then
		-- acks
		packet = packet .. string.char(0x00) .. -- ack
			Helpers.int_to_bytes(def.sequence_nr) -- seq nr

	end

	return packet
end

local function parse(buf)
	assert(#buf > 10, "packet too small")

	for i=1,4 do
		assert(string.byte(Constants.protocol_id, i) == string.byte(buf, i))
	end

	local def = {}

	def.peer_id = Helpers.bytes_to_int( string.byte(buf, 5), string.byte(buf, 6) )
	def.channel = string.byte(buf, 7)
	def.type = packet_types_by_id[string.byte(buf, 8)]
	def.sequence_nr = Helpers.bytes_to_int( string.byte(buf, 9), string.byte(buf, 10) )
	def.subtype = packet_types_by_id[string.byte(buf, 11)]
	def.payload = ""

	return def
end

-- exports
return {
	create = create,
	parse = parse
}
