package config

func init() {
	config := facades.Config
	config.Add("app", map[string]interface{}{
		//Application Name
		//This value is the name of your application. This value is used when the
		//framework needs to place the application's name in a notification or
		//any other location as required by the application or its packages.
		"name": config.Env("APP_NAME", "QuarkGo"),

		//Application Environment
		//This value determines the "environment" your application is currently
		//running in. This may determine how you prefer to configure various
		//services the application utilizes. Set this in your ".env" file.
		"env": config.Env("APP_ENV", "production"),

		//Application Debug Mode
		"debug": config.Env("APP_DEBUG", false),

		//Encryption Key
		//32 character string, otherwise these encrypted strings
		//will not be safe. Please do this before deploying an application!
		"key": config.Env("APP_KEY", ""),

		//Application host, http server listening address.
		"host": config.Env("APP_HOST", "127.0.0.1:3000"),
	})
}
