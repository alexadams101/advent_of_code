import scala.io.Source

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

def part1(lines: Iterator[String]): String = 
    val matrix = lines.map(_.toCharArray()).toArray
    var array = List[(Int, Boolean)]()

    var no = 0
    var adj = false

    for (y <- 0 until matrix.length) {
        for (x <- 0 until matrix(0).length) {
            if (matrix(y)(x).isDigit) {
                no = no * 10 + matrix(y)(x).toInt - 48
                if (hasAdjacentSymbol(matrix, y, x)) {
                    adj = true
                }
            } else if (no > 0) {
                array = array:+(no, adj)
                no = 0
                adj = false
            }
        }
    }

    array.filter(_._2).map(_._1).sum.toString()
    
def part2(lines: Iterator[String]): String = 
    val matrix = lines.map(_.toCharArray()).toArray
    var numbers = List[Int]()

    for (y <- 0 until matrix.length) {
        for (x <- 0 until matrix(0).length) {
            if (matrix(y)(x) == '*') {
                val n = getAdjacentNumbers(matrix, y, x)
                if (n.isDefined) {
                    numbers = numbers :+ n.get._1 * n.get._2
                }
            }
        }
    }

    numbers.sum.toString()

def getAdjacentNumbers(matrix: Array[Array[Char]], y: Int, x: Int): Option[(Int, Int)] =
    val coords = getCoordCombinations(y, x, matrix.length - 1, matrix(0).length - 1)
    var list = List[Int]()
    for (yI <- coords._1) {
        var c = 0
        for (xI <- coords._2) {
            if (matrix(yI)(xI).isDigit && xI >= c) {
                var i = xI - 1
                while (i >= 0 && matrix(yI)(i).isDigit) {
                    i = i - 1
                }   
                var n = 0
                i = i + 1
                while (i < matrix(yI).length && matrix(yI)(i).isDigit) {
                    n = n * 10 + (matrix(yI)(i).toInt - 48)
                    i = i + 1
                }
                list = list :+ n
                c = i
            }
        }
    }
    if (list.size == 2) {
        Some(list(0),list(1))
    } else {
        None
    }

def getCoordCombinations(y: Int, x: Int, yMax: Int, xMax: Int): (List[Int], List[Int]) = 
    var yL = List[Int]()
    var xL = List[Int]()

    yL = yL:+y
    xL = xL:+x

    if (y > 0) {
        yL = yL:+y-1
    }
    if (y < yMax) {
        yL = yL:+y+1
    }
    if (x > 0) {
        xL = xL:+x-1
    }
    if (x < xMax) {
        xL = xL:+x+1
    }

    (yL.sorted, xL.sorted)

def hasAdjacentSymbol(matrix: Array[Array[Char]], y: Int, x: Int): Boolean =
    val coords = getCoordCombinations(y, x, matrix.length - 1, matrix(0).length - 1)
    var output = false
    for (yI <- coords._1) {
        for (xI <- coords._2) {
            if (isSymbol(matrix(yI)(xI))) {
                output = true
            }
        }
    }
    output

def isSymbol(c: Char): Boolean = 
    !c.isDigit && c!='.'