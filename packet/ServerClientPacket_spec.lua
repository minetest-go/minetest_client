
local ServerClientPacket = require("packet/ServerClientPacket")

local function test_packet_definition(txdef)
	-- create -> parse
	local packet = ServerClientPacket.create(txdef)
	local def = ServerClientPacket.parse(packet)

	assert.are.same(txdef, def)
end

describe("ServerClientPacket test", function()
	it("reliable packet", function()
		-- TODO
	end)
end)
