import scala.io.Source

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

class Game(var instructions: String, var nodes: Map[String, Node]) {
    def getShortestTurnsPart1(): Int =
        def req(s: String): Boolean = s != "ZZZ"
        findTurns(nodes.get("AAA").get, req)

    def getShortestTurnsPart2(): BigInt = 
        val firstIds = nodes.keySet.filter(_.charAt(2) == 'A')
        var currentNodes = firstIds.map(nodes.get(_).get)
        var n = 0
        var count = 0
        def req(s: String): Boolean = s.charAt(2) != 'Z'
        BigDecimal(lcm(currentNodes.map(findTurns(_, req)))).toBigInt

    def findTurns(node: Node, req: String => Boolean): Int = 
        var currentNode = node
        var n = 0
        var count = 0
        while (req(currentNode.id))
            if (instructions.charAt(n) == 'L')
                currentNode = currentNode.left
            if (instructions.charAt(n) == 'R')
                currentNode = currentNode.right
            if (n == instructions.length() - 1)
                n = 0 
            else
                n = n + 1
            count = count + 1
        count

}

class Node(var id: String) {
    var left: Node = null
    var right: Node = null
}

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

def part1(lines: Iterator[String]): String = 
    val game = toGame(lines)
    game.getShortestTurnsPart1().toString()
    
def part2(lines: Iterator[String]): String = 
    val game = toGame(lines)
    game.getShortestTurnsPart2().toString()

def toGame(lines: Iterator[String]): Game =
    val instructions = lines.next()
    val nLines = lines.drop(1).toList
    val nodes = nLines.map(toNode).toMap
    val updatedNodes = nLines.map(updateNodes(nodes, _)).toMap
    new Game(instructions, updatedNodes)

def toNode(s: String): (String, Node) = 
    val id = s.substring(0, 3)
    (id, new Node(id))

def updateNodes(nodes: Map[String, Node], s: String): (String, Node) =
    val id = s.substring(0, 3)
    val left = s.substring(7, 10)
    val right = s.substring(12, 15)
    nodes.get(id).get.left = nodes.get(left).get
    nodes.get(id).get.right = nodes.get(right).get
    (id, nodes.get(id).get)

def lcm(s: Set[Int]): Double = 
    def count(i: Int, l: Set[Map[Int, Int]]): Double = 
        var max = 0
        for (m <- l)
            if (m.contains(i) && m.get(i).get > max)
                max = m.get(i).get
        Math.pow(i.toDouble, max.toDouble)

    val primes = s.map(factors).map(_.groupBy(identity).mapValues(_.size).toMap)
    val unique = primes.flatMap(_.keySet).toSet
    unique.map(count(_, primes)).product

def factors(n:Int):List[Int] = {
    def divides(d:Int, n:Int) = (n % d) == 0
    def ld(n:Int):Int =  ldf(2, n)
    def ldf(k:Int, n:Int):Int = {
      if (divides(k, n)) k
      else if ((k*k) > n) n
      else ldf((k+1), n)
    }
    n match {
      case 1 => Nil
      case _ => val p = ld(n); p :: factors(n / p)
    }
  }
