main :: IO ()
main = putStrLn "hello"

data Option a = Some a | None

safe_root :: Float -> Option Float
safe_root x = if x >= 0 then
        Some $ x / 2
    else
        None

safe_reciprocal :: Float -> Option Float
safe_reciprocal 0 = None
safe_reciprocal n = Some(1/n)

compose :: (a -> Option b) -> (b -> Option c) -> (a -> Option c)
compose f h x = case f x of
    Some(y) -> h y
    None -> None

(>=>) = compose

identity :: a -> Option a
identity = Some

f :: Float -> Option Float
f = safe_root >=> identity >=> safe_reciprocal
