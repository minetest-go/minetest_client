
local ClientServerPacket = require("packet/ClientServerPacket")

describe("ToServerPacket test", function()
	it("should parse the generated packet", function()
		local txdef = {
			-- command = "HELLO",
			peer_id = 0,
			channel = 0,
			type = "reliable",
			subtype = "original",
			sequence_nr = 65500,
			payload = ""
		}

		-- create -> parse
		local packet = ClientServerPacket.create(txdef)
		local def = ClientServerPacket.parse(packet)

		assert.are.same(txdef, def)
	end)
end)
