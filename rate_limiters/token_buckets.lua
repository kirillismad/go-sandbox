local key = KEYS[1]
local last_access_key = key .. ":last_access"
local now = tonumber(ARGV[1])
local unit = tonumber(ARGV[2])
local limit = tonumber(ARGV[3])
local ttl = tonumber(ARGV[4])

-- Получить текущее состояние
local tokens = tonumber(redis.call("GET", key) or limit)
local last_access = tonumber(redis.call("GET", last_access_key) or now)

-- Рассчитать пополнение
local time_passed = now - last_access
local rate = limit / unit
local replenished_tokens = math.floor(time_passed * rate)
tokens = math.min(tokens + replenished_tokens, limit)

if tokens < 1 then
    -- Недостаточно токенов
    return {0, now + math.ceil(1 / rate)}
end

-- Обновляем состояние
tokens = tokens - 1
redis.call("SET", key, tokens, "EX", ttl)
redis.call("SET", last_access_key, now, "EX", ttl)

return {1, tokens}
