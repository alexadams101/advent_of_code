import scala.io.Source
import scala.collection.immutable.Range
import scala.collection.immutable.NumericRange

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

class Ranging(var destination: Long, var source: Long, var length: Long)

class Almanac(var seeds: List[Long], var maps: List[List[Ranging]]):
    def calculateMinLocation(): Long =
        seeds.map(getLocation(_)).min
    
    def getLocation(seed: Long): Long = 
        var s = seed
        
        for (i <- 0 until maps.length) {
            s = map(s, maps(i))
        }
        s

    def calculateMinLocation(range: List[Long]): Long = 
        var l = Long.MaxValue
        for (i <- range(0) until range(0) + range(1)) {
            if (getLocation(i) < l) {
                l = getLocation(i)
            }
        } 
        l
    
    def map(seed: Long, ranges: List[Ranging]): Long = 
        val update = ranges.filter(validUpdate(seed, _))
        if (update.isEmpty) seed
        else seed - update(0).source + update(0).destination

    def map(l: List[Long]): NumericRange[Long] = 
        l(0) until l(0)+l(1)

    def calculateMinLocationPart2(): Long =
        println(seeds.grouped(2).mkString)
        seeds.grouped(2).map(calculateMinLocation).min

def validUpdate(seed: Long, r: Ranging): Boolean = 
    seed >= r.source && seed < r.source + r.length

def part1(lines: Iterator[String]): String = 
    val almanac = toAlmanac(lines)
    almanac.calculateMinLocation().toString()

def part2(lines: Iterator[String]): String = 
    val almanac = toAlmanac(lines)
    almanac.calculateMinLocationPart2().toString()

def toAlmanac(lines: Iterator[String]): Almanac =
    val seeds = lines.next().split(" ").drop(1).map(_.toLong).toList
    val m = getMapLines(lines).map(_.map(toRange))
    new Almanac(seeds, m)

def getMapLines(lines: Iterator[String]): List[List[String]] =
    var maps = List[List[String]]()
    var s = List[String]()
    for (l <- lines.drop(1)) {
        if (l.equals("")) {
            maps = maps:+ s
            s = List[String]()
        } else if (l.charAt(0).isDigit){
            s = s:+ l
        }
    }
    maps = maps:+ s
    maps

def toRange(line: String): Ranging =
    val nos = line.split(" ")
    new Ranging(nos(0).toLong, nos(1).toLong, nos(2).toLong)