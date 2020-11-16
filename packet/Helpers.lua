
local function int_to_bytes(i)
	local h = math.floor(i / 256) % 256;
	local l = math.floor(i % 256);
	return string.char(h, l);
end

local function bytes_to_int(h, l)
	return (h * 256) + l
end


return {
	int_to_bytes = int_to_bytes,
	bytes_to_int = bytes_to_int
}
