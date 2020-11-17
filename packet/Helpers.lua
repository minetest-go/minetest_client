
local function int_to_bytes(i)
	local h = math.floor(i / 256) % 256;
	local l = math.floor(i % 256);
	return string.char(h, l);
end

local function bytes_to_int(h, l)
	return (h * 256) + l
end

local function extract_string(payload, offset)
	local size = bytes_to_int( string.byte(payload, offset+1), string.byte(payload, offset+2) )
	return string.sub(payload, offset+3, offset + size + 3), size+2
end

local function create_string(str)
	return int_to_bytes(#str) .. str
end

return {
	int_to_bytes = int_to_bytes,
	bytes_to_int = bytes_to_int,
	extract_string = extract_string,
	create_string = create_string
}
