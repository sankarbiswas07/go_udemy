
SYNTAX:

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

A few important note:

1. Interfaces are not generic type.
   => Other language have 'generic' types - go (famously) does not.

2. Interface are implicit.
   => we can not put together any link between englishBOt & spanishBot.
   => we don't manually have to say that our custom type satisfies some interface

3. Interfaces are a contract to help us manage types.
   => Garbage in > Garbage out. If our custom type's implementation
   of a function is broken then interface won't hlp us!

4. Interfaces are tough. Step #1 is understanding how to read them.
   => Understand how to read interfaces in the standard lib.
   => Writing your own interface is tough and requires experience


=================================================



=================================================