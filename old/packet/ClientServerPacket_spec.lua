
local ClientServerPacket = require("packet/ClientServerPacket")

local function test_packet_definition(txdef)
	-- create -> parse
	local packet = ClientServerPacket.create(txdef)
	local def = ClientServerPacket.parse(packet)

	assert.are.same(txdef, def)
end

describe("ClientServerPacket test", function()
	it("reliable packet", function()
		test_packet_definition({
			peer_id = 0,
			channel = 0,
			type = "reliable",
			subtype = "original",
			sequence_nr = 65500,
			payload = ""
		})
	end)
end)
