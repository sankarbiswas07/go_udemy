syntax


type bot interface {
  getGreeting(string, int) (string, error)
}

here bot is interface name
getGreeting is the function name
(string, int) is the list of argument types
(string, error) is the list of return types

=================================================

type user struct{
  name string
}

type bot interface {
  getGreeting(string, int) (string, error)
  getBotVersion() float64
  respondToUser(user) string
}

=================================================
THESE TERMS WILL BE THERE IN GO LEARNING PATH

Concrete types : as I can create a value directly with 
=> map, struct, int, string, englishBot

Interface type : we can not actually create directly with
=> bot
=================================================
