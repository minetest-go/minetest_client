
local Helpers = require("packet/Helpers")

local function test_to_from_bytes(input)
	local bytes = Helpers.int_to_bytes(input)

	local high = string.byte(bytes, 1)
	local low = string.byte(bytes, 2)

	local output = Helpers.bytes_to_int(high, low)

	assert.are.same(input, output)
end

describe("Helpers test", function()
	it("should match high/low expectations", function()
		assert.are.same(65500, Helpers.bytes_to_int(0xff, 0xdc))
	end)

	it("should parse the generated int/byte", function()
		test_to_from_bytes(0)
		test_to_from_bytes(16000)
		test_to_from_bytes(40000)
		test_to_from_bytes(65000)
	end)
end)
