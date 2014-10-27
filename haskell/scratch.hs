import System.Environment

--main :: IO ()
--main = do
    --args <- getArgs
    --putStrLn $ evalRPN $ args !! 0

evalRPN :: String -> Float
evalRPN = head . foldl f [] . words
    where f (x:y:ys) "*" = (x * y):ys
          f (x:y:ys) "+" = (x + y):ys
          f (x:y:ys) "-" = (x - y):ys
          f (x:y:ys) "/" = (x / y):ys
          f (x:y:ys) "^" = (x ** y):ys
          f (x:xs) "ln" = log x:xs
          f xs "sum" = [sum xs]
          f xs num = read num:xs
