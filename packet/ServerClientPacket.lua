local Constants = require("packet/Constants")
local Helpers = require("packet/Helpers")
local PacketTypes = require("packet/PacketTypes")
local ControlTypes = require("packet/ControlTypes")
local ServerClientCommands = require("packet/ServerClientCommands")

-- lookup tables
local commands_by_id = {}
local commands_by_key = {}

for _, cmd in ipairs(ServerClientCommands) do
	commands_by_id[cmd.id] = cmd
	commands_by_key[cmd.key] = cmd
end

local control_types_by_id = {}
for k, id in pairs(ControlTypes) do
	control_types_by_id[id] = k
end

local packet_types_by_id = {}
for k, id in pairs(PacketTypes) do
	packet_types_by_id[id] = k
end


local function parse(buf)
	assert(#buf >= 9, "packet too small")

	for i=1,4 do
		assert(string.byte(Constants.protocol_id, i) == string.byte(buf, i))
	end

	-- 4F457403 0001 00 03 FFDC 00 01 0141
	-- 4F457403 0001 00 03 FFDC 00 01 00F4
	-- 4F457403 0001 00 00 00 FFDC

	--                  __ < reliable
	--                          __ < subtype
	-- 4F457403 0001 00 03 FFDD 00 02
	-- 4F457403 0001 00 03 FFDD 01 00 02 1C00000027000000020004626C6168
	-- 4F457403 0001 00 03 FFDD 01 00 02 1C00000027000000020004626C6168
	-- 4F457403 0001 00 03 FFDD 01 00 02 1C00000027000000020004626C6168

	-- 4F457403 0001 00 03 FFDE 01 00 29 3F5C42900000
	-- 4F457403 0001 00 03 FFDF 00 02
	-- 4F457403 0001 00 03 FFDC 00 01 015B -- set peer id
	-- 4F457403 0001 00 03 FFDC 00 01 0187
	-- 4F457403 0001 00 03 FFDC 00 01 0198
	-- 4F457403 0001 00 03 FFDD 00 02

	-- 4F457403 0001 00 00 03

	local def = {}

	def.peer_id = Helpers.bytes_to_int( string.byte(buf, 5), string.byte(buf, 6) )
	def.channel = string.byte(buf, 7)

	def.type = packet_types_by_id[string.byte(buf, 8)]

	if def.type == "reliable" then
		def.sequence_nr = Helpers.bytes_to_int( string.byte(buf, 9), string.byte(buf, 10) )
		def.subtype = packet_types_by_id[string.byte(buf, 11)]

		if def.subtype == "control" then
			-- control packet
			def.controltype = control_types_by_id[string.byte(buf, 12)]

			if def.controltype == "SET_PEER_ID" then
				-- set peer id
				local cmd = commands_by_key["SET_PEER_ID"]
				if not cmd then
					error("unknown command received" .. def.command_id)
				end

				def.command = cmd.key
				def.payload = cmd.parse(buf:sub(13))

			end

		elseif def.subtype == "original" then
			-- normal packet with subtype
			def.command_id = Helpers.bytes_to_int( string.byte(buf, 12), string.byte(buf, 13) )

			local cmd = commands_by_id[def.command_id]
			if not cmd then
				error("unknown command received: " .. def.command_id)
			end

			def.command = cmd.key
			def.payload = cmd.parse(buf:sub(14))
		else
			error("unknown subtype")
		end

	elseif def.type == "original" then
		error("not implemented")

	elseif def.type == "control" then
		def.controltype = control_types_by_id[string.byte(buf, 9)]

		if def.controltype == "ACK" then
			-- ack
			def.sequence_nr = Helpers.bytes_to_int( string.byte(buf, 10), string.byte(buf, 11) )

		elseif def.controltype == "DISCO" then
			error("disconnect")

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
