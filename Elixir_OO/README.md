# final-exam

# final-exam

Explain what features in Elixir you are using to provide the objects and OO pillars you are implementing. You may or may not have all of them... so explain the ones you do have in your code solution. Explain what parts of your solution map match what part of the Java example you are implementing.


The OO pillars are encapsulation, abstraction, polymorphism, and inheritance.

For encapsulation in the "Animal" Java class, I make the name private to the scope of the animal process in Elixir. The name is passed into the animal class when the message is sent in the "new" function. The name value is only accessible inside the animal process. Another process like "dog" and "cat" cannot get the name value from the "animal" class. In Elixir, I don't have a getName message, but I use spawn and link to create a dog/cat process with the name as the parameter value. This step allows the dog/cat process to have access to the name variable indirectly. 

For abstraction in the "Animal" Java class, I made something similar to the animal class in Elixir. Inside the body of the "{:speak, _pid}" message pattern matching, I didn't implement the print statement directly. Instead, I pass the specific implementation to the "dog" and "cat" processes by just spawning those processes in the body of "{:speak, _pid}". This abstract the speak function in the "animal" process. 

Polymorphism in the Java example has two versions of the speak method body. In Elixir, both the "dog" and "cat" processes have the ":speak" message pattern matching. The speak method in dog and cat have different print outputs, which satisfy the principle of polymorphism. 

For inheritance in the Java example, both the dog and cat classes inherit from the animal class. In Elixir, the name of either dog or cat is accessible within the process but only can be passed down as a parameter from the animal process. Both the dog and cat process have the ":speak" message pattern matching, which is designed as inherited from the animal process's speak message pattern matching as well. 




The following explains the parts of my solution that map to the Java example. The Elixir new function matches the class Main, where two animal objects are created and each calls the speak function. In Elixir, the two animals are spawned and linked to the animal process with their name ("Rex", "Whiskers") and their type ("dog", "cat"). The speech message is then passed to the animal class for each animal object. 

The "animal" process in Elixir matches the "Animal" class in Java. In Java, the animal class has the constructor set the name,  one "getName" function to get the name, and one "speak" abstract function. In the animal class in Elixir, the name and animal type are passed in as parameters so the name for an animal object is set when a message is sent from "new". In the body of the "{:speak, _pid}" pattern matching, there is no direct implementation, only spawning the corresponding animal type using the name. 

The "dog" process in Elixir matches the "Dog" class in Java. The "cat" process in Elixir matches the "Cat" class in Java. In Java, the "Dog" class constructor uses the name from its parent class "Animal" and it has an implementation for the speak function, which prints the dog's name and "says: Woof!". Similarly for the "Cat" class, the only difference is its print statement being the cat's name and "says: Meow!". 




