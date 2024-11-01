#include <iostream>
using namespace std;

class Singleton
{
private:
    static Singleton *instancePtr;
    Singleton()
    {
        cout << "singleton created once" << endl;
    }

public:
    Singleton(const Singleton &) = delete;
    static Singleton *getInstance()
    {
        return instancePtr;
    }
    void show()
    {
        cout << "hello from singleton class" << endl;
    }
};

Singleton *Singleton::instancePtr = new Singleton();

int main()
{
    Singleton *instance1 = Singleton::getInstance();
    instance1->show();

    // Trying to get another instance just gives us the same one
    Singleton *instance2 = Singleton::getInstance();
    instance2->show();

    // Showing that both pointers point to the same instance
    cout << "Instance1 address: " << instance1 << endl;
    cout << "Instance2 address: " << instance2 << endl;
    return 0;
}