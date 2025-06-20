module Main where

import qualified Data.Map as Map
import qualified Data.List as List
import System.Directory as Dir
import qualified Distribution.Compat.Directory as Dir

a :: DataBase
a = DataBase (Map.fromList [("key1", Frequency [("file1", 2), ("file2", 1)])])

b :: DataBase
b = DataBase (Map.fromList [("key1", Frequency [("file1", 3), ("file3", 3)])])

fq1 = DataBase $ Map.fromList $ wrodsFrequency ["hello", "there", "hello"] "file1"
fq2 = DataBase $ Map.fromList $ wrodsFrequency ["hello", "there", "hello"] "file2"

main = do
    print a
    print b
    print $ a <> b
    print $ wrodsFrequency ["hello", "there", "hello"] "file"
    print $ List.group ["hello", "there", "hello"] 
    print $ List.group [1,2,3,1] 
    print $ "hello" == "hello"
    print $ fq1 <> fq2

newtype DataBase = DataBase (Map.Map KeyWord Frequency) deriving (Show)
type KeyWord = String
newtype Frequency = Frequency [(FileName, Count)] deriving (Show)
type FileName = String
type Count = Int

union :: Frequency -> Frequency -> Frequency
union (Frequency a) (Frequency b) = Frequency $ List.sortBy (\(_, m) (_, n) -> compare n m) $ Map.toList $ Map.unionWith (+) (Map.fromList a) (Map.fromList b)

wrodsFrequency :: [String] -> FileName -> [(KeyWord, Frequency)]
wrodsFrequency xs fn = (\(kw, count) -> (kw, Frequency [(fn, count)])) <$> Map.toList (Map.fromListWith (+) [(x, 1::Count) | x <- xs])

instance Semigroup DataBase where
    (DataBase a) <> (DataBase b) = DataBase (Map.unionWith union a b)

instance Monoid DataBase where
    mempty = DataBase (Map.empty :: Map.Map KeyWord Frequency)

