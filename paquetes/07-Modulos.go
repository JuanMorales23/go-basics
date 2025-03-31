//Son una herramienta crucial para gestionar las dependencias de un proyecto
//Go modules es una herramienta que se utiliza para administrar las dependencias, permite escribir y especificar
//Las versiones de las dependencias externas y garantiza la integridad de las mismas
//Permite gestionar las versiones tanto de dependencias externas como de dependencias internas
//Para inicializar un módulo se utiliza el comando go mod init nombreDelModulo
//Para agregar una dependencia se utiliza el comando go get nombreDelModulo
//Para actualizar una dependencia se utiliza el comando go get -u nombreDelModulo
//Para eliminar una dependencia se utiliza el comando go mod tidy
//Para ver las dependencias se utiliza el comando go list -m all
//Para ver las dependencias de un módulo en específico se utiliza el comando go list -m -json nombreDelModulo
//En el archivo go.mod se pueden ver las dependencias del proyecto, este debe estar situado en la raíz del proyecto

//Antes de inciar un módulo tenemos que crear un nuevo repositorio de github u otroso controladores de versiones
//go mod init (urlGithub)
//go get (urlGithub)
//con el flag -u permite actualizar las dependencias
//Luego de agregar la dependencia se puede agregar el módulo dentro de nuestros archivos usando el import