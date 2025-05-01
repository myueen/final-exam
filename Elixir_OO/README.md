# final-exam

Explain what features in Elixir you are using to provide the objects and OO pillars you are implementing. You may or may not have all of them... so explain the ones you do have in your code solution. Explain what parts of your solution maps to what part of the Java example you are implementing.


The OO pillars are encapsulation, abstraction, polymorphism, and inheritance.

For encapsulation in the "Animal" Java class, we make the name private to the scope of the animal process in Elixir. The name is passed into the animal class when the message is sent in the "new" function. The name value is only accessible inside the animal process. Other process like "dog" and "cat" cannot get the name value from the "animal" class. In Elixir, we don't have a getName message, but we use spawn and link to create a dog/cat process with the name as parameter value. This stepss allows the dog/cat process to have access to the name variable indirectly. 

For abstraction in the "Animal" Java class, we make something similar in the animal class in Elixir. Inside the body of the "{:speak, _pid}" message pattern matching, we didn't implement the print statement directly. Instead, we pass the specific implementation to the "dog" and "cat" process by just spawning those processes in the body of "{:speak, _pid}". This abstract the speak function in the "animal" process. 

For polymorphism in Java example, they have two versions of the speak method body. In Elixir, both "dog" and "cat" process have the ":speak" message pattern matching that's 





The following explain the parts of my solution that maps to the Java example. The Elixir new function matches the class Main, where two animals objects are created and each calls the speak function. In Elixir, the two animals are spawned and link to the animal process with their name ("Rex", "Whiskers") and their type ("dog", "cat"). The speak message is then passed to the animal class for each animal object. 

The "animal" process in Elixir match the "Animal" class in Java. In Java, the animal class have the constructor sets the name,  one "getName" function to get the name, and one "speak" abstract function. In the animal class in Elixir, the name and animal type is passed in as parameters so the name for an animal object is set when a message is sent from "new". In the body of the "{:speak, _pid}" pattern matching, there is no direct implementation, only spawning the corresponding animal type using the name. 

The "dog" process in Elixir matches the "Dog" class in Java. The "cat" process in Elixir matches the "Cat" class in Java. In Java, the "Dog" class constructor uses got the name from its parent class "Animal" and it has implementation for the speak function, which prints the dog's name and "says: Woof!". Similarly for the "Cat" class with the only different is its print statement being the cat's name and "says: Meow!". 




