# final-exam

Explain what features in Elixir you are using to provide the objects and OO pillars you are implementing. You may or may not have all of them... so explain the ones you do have in your code solution. Explain what parts of your solution maps to what part of the Java example you are implementing.


The OO pillars are encapsulation, abstraction, polymorphism, and inheritance.





The following explain the parts of my solution that maps to the Java example. The Elixir new function matches the class Main, where two animals objects are created and each calls the speak function. In Elixir, the two animals are spawned and link to the animal process with their name ("Rex", "Whiskers") and their type ("dog", "cat"). The speak message is then passed to the animal class for each animal object. 

The "animal" process in Elixir match the "Animal" class in Java. In Java, the animal class have the constructor sets the name,  one "getName" function to get the name, and one "speak" abstract function. In the animal class in Elixir, the name and animal type is passed in as parameters so the name for an animal object is set when a message is sent from "new". In the body of the "{:speak, _pid}" pattern matching, there is no direct implementation, only spawning the corresponding animal type using the name. 

The "dog" process in Elixir matches the "Dog" class in Java. The "cat" process in Elixir matches the "Cat" class in Java. In Java, the "Dog" class constructor uses got the name from its parent class "Animal" and it has implementation for the speak function, which prints the dog's name and "says: Woof!". Similarly for the "Cat" class with the only different is its print statement being the cat's name and "says: Meow!". 




