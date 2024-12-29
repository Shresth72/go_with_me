package main

func main() {
  // --- Client
  // Initial Requirement
  apple := &apple{}
  client := &client{}
  client.ChargeMobile(apple)

  // Extended Requirement
  android := &android{}
  androidAdapter := &androidAdapter{
    android: android,
  }
  client.ChargeMobile(androidAdapter)

  // --- Chess
  knight := &thirdPartyKnight{}
  bishop := &thirdPartyBishop{}

  knightAdapter := &knightAdapter{knight: knight}
  bishopAdapter := &bishopAdapter{bishop: bishop}

  player := &Player{}
  player.PlayMove(knightAdapter)
  player.PlayMove(bishopAdapter)
}
