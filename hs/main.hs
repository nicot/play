main :: IO ()
main = print "foo"

data List a f = Cons a (List a f) | Nil

foo :: List Integer f
foo = Cons 1 (Cons 3 Nil)

newtype Fix f = In (f (Fix f))

(|>) :: a -> (a -> b) -> b
(|>) x f = f x

zah :: String -> Int
zah s = s |> length

k :: Fix (List Integer)
k = In (Cons 1 Nil)

fix :: (a -> a) -> a -> a
fix f = fix f . f

sum' :: Integer -> Integer -> Integer
sum' x y = x + y

cata :: (b -> a -> b) -> Fix (f a) -> b
cata g f = undefined

that :: Integer
that = cata sum' k