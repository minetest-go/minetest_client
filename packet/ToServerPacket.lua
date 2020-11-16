
local Constants = require("packet/Constants")
local Helpers = require("packet/Helpers")
local PacketType = require("packet/PacketType")

local commands = {
	["HELLO"] = "\x00\x01"
}

local function create(def)
	-- assert(commands[def.command], "command not available: " .. def.command)

	local packet = 	Constants.protocol_id ..
		Helpers.int_to_bytes(def.peer_id) .. -- peer_id
		string.char(def.channel) .. -- channel
		PacketType[def.type] .. -- type
		Helpers.int_to_bytes(def.sequence_nr) .. -- seq nr
		PacketType[def.subtype] .. -- subtype
		def.payload

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

	for k, v in pairs(PacketType) do
		if string.byte(v) == string.byte(buf, 8) then
			def.type = k
			break
		end
	end

	def.sequence_nr = Helpers.bytes_to_int( string.byte(buf, 9), string.byte(buf, 10) )

	for k, v in pairs(PacketType) do
		if string.byte(v) == string.byte(buf, 11) then
			def.subtype = k
			break
		end
	end

	def.payload = ""

	return def
end

-- exports
return {
	create = create,
	parse = parse
}
