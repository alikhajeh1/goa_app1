package design

import (
        . "github.com/goadesign/goa/design"
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa_app1", func() {
        Title("Example app 1")
        Description("A simple goa service")
        Scheme("http")
        Host("localhost:8080")
})

var _ = Resource("numbers", func() {
        BasePath("/numbers")
        DefaultMedia(NumberMedia)

        Action("show", func() {
                Description("Get prime numbers")
                Routing(GET("/"))
                Params(func() {
                        Param("lessThan", Integer, "Prime numbers less than this number are calculated")
                })
                Response(OK)
        })
})

var NumberMedia = MediaType("application/vnd.goa.example.numbers+json", func() {
        Description("A list of number")
        Attributes(func() {
                Attribute("numbers", ArrayOf(Integer), "Numbers")
        })
        View("default", func() {
                Attribute("numbers")
        })
})

var _ = Resource("swagger", func() {
        Origin("*", func() {
               Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
        })
        Files("/swagger.json", "swagger/swagger.json")
})