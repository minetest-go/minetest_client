local Helpers = require("packet/Helpers")

return {
	{
		id = 2,
		key = "INIT",
		create = function(def)
			return string.char(def.client_max) ..
				Helpers.int_to_bytes(def.supp_compr_modes) ..
				Helpers.int_to_bytes(def.min_net_proto_version) ..
				Helpers.int_to_bytes(def.max_net_proto_version) ..
				string.char(0x00) .. string.char(#def.player_name) ..
				def.player_name
		end
	}
}
