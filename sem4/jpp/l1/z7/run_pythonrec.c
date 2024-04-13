#include <Python.h>

int main() {
    Py_Initialize();

    PyRun_SimpleString("import sys");
    PyRun_SimpleString("sys.path.append('.')");

    PyObject* module_name = PyUnicode_DecodeFSDefault("mymathrec");
    PyObject* module = PyImport_Import(module_name);
    Py_DECREF(module_name);
    Py_DECREF(module_name);
    if (!module) {
        PyErr_Print();
        printf("Failed to load mymathloop.py\n");
        return 1;
    }

    PyObject *diophantine = PyObject_GetAttrString(module, "diophantine");
    if (!diophantine || !PyCallable_Check(diophantine)) {
        PyErr_Print();
        printf("No such function");
        return 1;
    }

    PyObject* diophantine_args = PyTuple_Pack(3, PyLong_FromLong(60), PyLong_FromLong(94), PyLong_FromLong(30));
    PyObject* diophantine_results = PyObject_CallObject(diophantine, diophantine_args);

    if (diophantine_results != NULL) {
        printf("Result of diophantine(60, 94, 23): (%ld, %ld)\n", PyLong_AsLong(PyTuple_GetItem(diophantine_results, 0)), PyLong_AsLong(PyTuple_GetItem(diophantine_results, 1)));
        Py_DECREF(diophantine_results);
    } else {
        PyErr_Print();
        printf("Call failed\n");
    }

    // Function 'factorial' (silnia)
    PyObject *factorial = PyObject_GetAttrString(module, "factorial");
    if (!factorial || !PyCallable_Check(factorial)) {
        PyErr_Print();
        printf("No such function 'factorial'\n");
        Py_DECREF(module);
        Py_Finalize();
        return 1;
    }

    PyObject* factorial_args = PyTuple_Pack(1, PyLong_FromLong(8));
    PyObject* factorial_result = PyObject_CallObject(factorial, factorial_args);
    Py_DECREF(factorial_args);

    if (factorial_result != NULL) {
        printf("Result of factorial(8): %ld\n", PyLong_AsLong(factorial_result));
        Py_DECREF(factorial_result);
    } else {
        PyErr_Print();
        printf("Call to 'factorial' failed\n");
    }
    Py_DECREF(factorial);

    // Function 'gcd' (nwd)
    PyObject *gcd = PyObject_GetAttrString(module, "gcd");
    if (!gcd || !PyCallable_Check(gcd)) {
        PyErr_Print();
        printf("No such function 'gcd'\n");
        Py_DECREF(module);
        Py_Finalize();
        return 1;
    }

    PyObject* gcd_args = PyTuple_Pack(2, PyLong_FromLong(20), PyLong_FromLong(14));
    PyObject* gcd_result = PyObject_CallObject(gcd, gcd_args);
    Py_DECREF(gcd_args);

    if (gcd_result != NULL) {
        printf("Result of gcd(20, 14): %ld\n", PyLong_AsLong(gcd_result));
        Py_DECREF(gcd_result);
    } else {
        PyErr_Print();
        printf("Call to 'gcd' failed\n");
    }
    Py_DECREF(gcd);

    // Clean up
    Py_DECREF(module);
    Py_DECREF(diophantine);
    Py_DECREF(diophantine_args);
    Py_Finalize();
    return 0;
}