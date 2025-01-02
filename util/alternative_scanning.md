```
func main() {
  input, _ := os.ReadFile("../input.txt")
  lines := strings.Split(strings.TrimSpace(string(input)), "\n")
  totalTokens := 0
  for i := 0; i < len(lines); i += 4 {
    var aX, aY, bX, bY, prizeX, prizeY int
    fmt.Sscanf(lines[i+0], "Button A: X+%d, Y+%d", &aX, &aY)
    fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bX, &bY)
    fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &prizeX, &prizeY)
    prizeX, prizeY = prizeX+10000000000000, prizeY+10000000000000
    D, Dx, Dy := aX*bY-bX*aY, prizeX*bY-bX*prizeY, aX*prizeY-prizeX*aY
    if D != 0 && Dx == (Dx/D)*D && Dy == (Dy/D)*D {
      totalTokens += (Dx/D)*3 + (Dy / D)
    }
  }
  fmt.Println(totalTokens)
}
```