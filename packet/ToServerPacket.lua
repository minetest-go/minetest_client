
local Constants = require("packet/Constants")
local PacketType = require("packet/PacketType")

local commands = {
	["HELLO"] = "\x00\x01"
}

local function create(def)
	assert(commands[def.command], "command not available: " .. def.command)

	local packet = 	Constants.protocol_id ..
		string.char(0x00, 0x00) .. -- peer_id
		string.char(0x00) .. -- channel
		PacketType.reliable .. -- type
		string.char(0xff, 0xdc) .. -- seq nr
		PacketType.original .. -- subtype
		string.char(0x00, 0x00)

	return packet
end

local function parse(buf)
	assert(#buf > 10, "packet too small")

	for i=1,4 do
		assert(string.byte(Constants.protocol_id, i) == string.byte(buf, i))
	end

	return {}
end

-- exports
return {
	create = create,
	parse = parse
}
