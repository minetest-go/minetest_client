
local ToServerPacket = require("packet/ToServerPacket")

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
		local packet = ToServerPacket.create(txdef)
		local def = ToServerPacket.parse(packet)

		assert.are.same(txdef, def)
	end)
end)
