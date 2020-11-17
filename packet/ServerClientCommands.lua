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
		parse = function(payload)
			return {
				serialization_ver = string.byte(payload, 1),
				compression_mode = Helpers.bytes_to_int( string.byte(payload, 2), string.byte(payload, 3) ),
				proto_ver = Helpers.bytes_to_int( string.byte(payload, 4), string.byte(payload, 5) )
				-- auth_mechs: 4 bytes,
				-- username_legacy: ?
			}
		end
	},
	{
		id = 41,
		key = "TIME_OF_DAY",
		parse = function(payload)
			return {
				time_of_day = Helpers.bytes_to_int( string.byte(payload, 1), string.byte(payload, 2) ) % 24000
			}
		end
	}
}
