local tokens_key = KEYS[1]
local timestamp_key = tokens_key .. ":timestamp"
local max_tokens = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = 1

-- Получить текущее состояние
local tokens = tonumber(redis.call("GET", tokens_key) or max_tokens)
local last_refill = tonumber(redis.call("GET", timestamp_key) or 0)

-- Рассчитать пополнение
local time_passed = now - last_refill
local tokens_to_add = math.floor(time_passed * refill_rate)
tokens = math.min(tokens + tokens_to_add, max_tokens)

if tokens < requested then
    -- Недостаточно токенов
    return {0, tokens}
end

-- Обновляем состояние
tokens = tokens - requested
redis.call("SET", tokens_key, tokens, "EX", 60 * 60 * 24) -- Установить срок хранения (1 день)
redis.call("SET", timestamp_key, now, "EX", 60 * 60 * 24)

return {1, tokens}
