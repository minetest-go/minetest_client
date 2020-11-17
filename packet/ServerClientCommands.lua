local Helpers = require("packet/Helpers")

return {
	{
		id = 1,
		key = "SET_PEER_ID",
		parse = function(payload)
			return {
				peer_id = Helpers.bytes_to_int( string.byte(payload, 1), string.byte(payload, 2) )
			}
		end
	},
	{
		id = 2,
		key = "HELLO",
		parse = function() end
	}
}
