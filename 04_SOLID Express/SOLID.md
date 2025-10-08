# SOLID
é um acrônimo que consolida 5 itens que são considerados como boas práticas no mundo do desenvolvimento orientado a objetos.

# 1 - SRP = Single Responsiblity Principle
Uma classe deve ter uma responsabilidade

# OCP = Open-closed Principle
A Classe após ser aberta, ela precisa ser fechada !

# LSP = Liskov Substitution Principle
As subclasses devem poder substituir a classe pai sem alterar o comportamento do programa

# Exemplo de Código

```python
class Animal:
    def rosnar(self):
        pass

class Cachorro(Animal):
    def rosnar(self):
        return "Latido"

-- Aqui fica errado, gato não rosna
class Gato(Animal):
    def rosnar(self):
        return "Miau"

def fazer_animal_rosnar(animal: Animal):
    print(animal.rosnar())

cachorro = Cachorro()
gato = Gato()

fazer_animal_rosnar(cachorro)
fazer_animal_rosnar(gato)
```

Este exemplo demonstra o Princípio da Substituição de Liskov (LSP), onde as subclasses `Cachorro` e `Gato` podem substituir a classe pai `Animal` sem alterar o comportamento do programa.


# ISP = Interface Segregation Principle
Uma classe não é obrigada a implementar os métodos da interfaces que ela não utilizará .
Há situações que os programadores deixam o método vazio ou coloca um  throw !


# DIP = Dependency Inversion Principle
Depender de abstrações.
Não ter acoplamento

