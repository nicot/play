import Text.ParserCombinators.Parsec hiding (spaces)
import System.Environment
import Control.Monad

main :: IO ()
main = do
    args <- getArgs
    putStrLn $ readExpr $ args !! 0

data LispVal = Atom String
    | List [LispVal]
    | DottedList [LispVal] LispVal
    | Number Integer
    | String String
    | Bool Bool

parseExpr :: Parser LispVal
parseExpr = parseAtom
    <|> parseString
    <|> parseNumber
    <|> parseQuoted
    <|> do char '('
           x <- try parseList <|> parseDottedList
           char ')'
           return x

parseQuoted :: Parser LispVal
parseQuoted = do
    char '\''
    x <- parseExpr
    return $ List [Atom "quote", x]

parseString :: Parser LispVal
parseString = do
    char '"'
    x <- many (noneOf "\"")
        <|> string "\"\""
        <|> string "\n"
    char '"'
    return $ String x

parseList :: Parser LispVal
parseList = liftM List $ sepBy parseExpr spaces

parseDottedList :: Parser LispVal
parseDottedList = do
    head <- endBy parseExpr spaces
    tail <- char '.' >> spaces >> parseExpr
    return $ DottedList head tail

parseAtom :: Parser LispVal
parseAtom = do
    first <- letter <|> symbol
    rest <- many (letter <|> digit <|> symbol)
    let atom = first:rest
    return $ case atom of
        "#t" -> Bool True
        "#f" -> Bool False
        _    -> Atom atom

parseNumber :: Parser LispVal
parseNumber = (many1 digit) >>=  (\n -> return $ Number $ read n)

symbol :: Parser Char
symbol = oneOf "!#$%&|*+-/:<=>?@^_~"

readExpr :: String -> String
readExpr input = case parse parseExpr "lisp" input of
    Left err -> "No match: " ++ show err
    Right _ -> "Found value"

spaces :: Parser ()
spaces = skipMany1 space
