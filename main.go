package main

import (
	"fmt"
	"gpt-project/router"
	"net/http"
	"os"
	"time"
)

func main() {
	//fmt.Println("Starting a service!!!!")
	route := router.SetupRouter()
	
	port := os.Getenv("PORT")
	ip := os.Getenv("IP")
	if ip == ""{
		ip = "localhost"
	}
	
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting a service on port!!!!"+port)
	//route.Use(corsMiddleware())
	s := &http.Server{
		Addr:           ip +":"+ port,
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start a server!!!")
		return
	}
}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		c.Next()
// 	}
// }

//
//import (
//	"context"
//	_ "database/sql/driver"
//	"fmt"
//	"github.com/nekomeowww/go-pinecone"
//	"github.com/tmc/langchaingo/llms/openai"
//	"log"
//)
//
//func main() {
//	//url := "https://vsgpt-gqeyhmw.svc.gcp-starter.pinecone.io/describe_index_stats"
//	////url := "https://index_name-project_id.svc.environment.pinecone.io/describe_index_stats"
//	//
//	//req, _ := http.NewRequest("POST", url, nil)
//	//
//	//req.Header.Add("accept", "application/json")
//	//req.Header.Add("content-type", "application/json")
//	//req.Header.Add("Api-Key", "3109588f-dbae-4e4f-84cb-e71ea6e60816")
//	//
//	////req.Header.Add("Api-Key", "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwiaXNzIjoiaHR0cHM6Ly9sb2dpbi5waW5lY29uZS5pby8ifQ..O3A9MU0OtrvXG2Fa.ALKQhDIjYE3MrkGey4ruYISF--hQbIlGzpoiDH9ARDkLhB7ZuHsokShIN2NfMwvghuv_N-Gv53LfBWhm5y1Mn1SqfPns5SBGkd3hougxPX3uQaqEXL9kwlL9ryFBwnHneGvrTIQsCL7zxzT2L6-UU4fBZN0TP5CyTnX7eIz39elQyWl-aDXM75ZJ3UpEiXiEbVaZe4VJvxJ7FDxPxURExqVzsym92ut6eWT6qnj-ckmLWWVWDqq5SP_p8VlPdfAjPA-BalDx7rW736WaR2FoHqaP9loGnntKRIma_cpss17fzCEwSKHINkNzPlsGPMDlaWGFe7Ef-HY9CA4xf5yfqgFi4X4bcJA0BGdRA2MaohPgbbM0IB0BSCfJHTDlutHF_DGDzDAD8_L2YZzmsWMkOY11u2Xs0QGawTclD4qL-aCZnLDkzc-Kwp3bmxbHCUdYKDLQBen5XA-C-Nde2MeNS_XosWLv8XESTO7OckvPBv4qUJ1hl_ml8MeJjCt9jskegB5IaY-SMcjEBYGCrSVlZOnbqYwOJaOuxMmkLUCjKzv4biE16E2CW7Milu8SgxQaK76E1_VZt67HFTlohKhDls91O_DeKyI0M6fAc1nrObFtcw0NulThPf302lQa6jKN4rygPALtA1srZT3JNpf15S_C_zEBJmd_Ue8Xlr9RWXRhpt6jQ1xt-G0L3RB6NwVSvTbGUsR-oitvH0sA9tNxIZjWL83eBVm6i9z32kvJEr_sqibZu3t0qKNkpcz6qqyzRKj7cE-9za9gJwLqMOe90cf9rWz3yNlgng0pCjXzIwfKv2MLPhe5iKIUdbDRxg61VEVaqyWhEUfVsLKRu6nxZhI80EgOnVICYMehMcQwWeXYvX3wG2NDTx3UsqiWyZ2imRPSK9fa3FKV6boe-6gElZOfo5bKlrFKaYWmMSw5YETfwvxtVdY46HsmmJ2nnYwvp2kX3OX1VUoEKunf-HvIjC5Ez3J6HIlNTzv31WtNUSS724AgqazR_bVdUm4KeXGn8hCQs4wzfIBlom6DI2ufGex3OwbCp941enELT-2jiB_4msN3e04Crnfqv-ay204G3J2asgdrVI9Dx6-PhOyoVK6tI4tQjZKiJhNN3r59DlO5ZpVzE7kgGJjFXTQXt2_u0qG1ttgTZdeiyIVF3gdNOjgtu-H6e3nTyGUIFklT2QjmQPxe78jnOOnnqzMehNlk6onoUJc2o3-1OFpCFNqZSTJ3MIh0579Skuz5eTw4vjUN040KoYWyetaqrLuyXWzXCZZCgwWVaIYWBuTI9PNE89dKsUpvL6fItEJvLFx4QwLep_lFRzqnTvUQPJtzc5LMvp4WVOxU5q1_EVGW4YjMh5rBNVK6N-cdACvAYjSTmNb0cGWNo_pS0dNKPK68NNQYPH_hFyDSxSFHwxL3h-lCb4rXxmJjlMDAB9C21bqbgqKigjZJ7-jC2eotr1LhbYWStZdOMn6k5AcE48AfJZAdaItdp5JEmuSv2Ezp6KsprYNOyacg2KD8H26lm_dACvfq8Nu44DY6zwvySx5RO4LW72_RZGU5O_4DOt_NqJrVEawz8jHGb5wA3GBaOGSLgNHsIeNFuTbQWS4st2sSctLHqVjuQ4OGGfxHjpohE5JFOFAwR0huTBbn3Yl3ic5WAajOLnhf2Eta1ivzlVVdhgFzSsKUiAJBGv6Rkw9hiIUFLpyRPDOdamBKT9IwpAwm3kpXGC6ZCJQCIOe_1JpIHglnifm5Jnl876MgCDM0P9QrmNyAwg1WUYCLfR4qUPdCrbz3sv5NgBlos8KApa_2P9W_zzkn1nfVEs6SMYwkq2LZEp-IstFObv_rdHepMmk61IGGHD0_tavBq3dQNYbZazqEvHvmuVLkOPOf4oQXb_VyRaW1npNv5NIkr-p_52BETDQB5PAdmYDyIDfWooVvugC-WjnQgvSiUG8w3j6mkcHxI9yrqDWKG_fGrRUo6sQmGwYalP60MTo7UIszVRirYPptx6iXGcox3XXXlPjFUGPC7-uLxLE222NcMKXmwz6asPnR2lu43A486kF5g6mic5GeiCNICzXN76BZNnvXBRWuxvjzlvI4tckAwWrYuDA71WCjzA0KDCZDCFi-Ji29Tsc14DonTr7C0ggCOctgWANXub2n1errfDz5V0b4n82dYzv_78c3Uz1oc8962gT5QwhTdYs4Z3tzaVE7GwYPhSk2EnXHc3hSIeU7v2LHITNn-CU1cbjlbEO7l6KsKhPRJm0Sulk0mgM8kwElauMvLT8Ulpv1Knt6oPX_IGXnp8Z5uBV192MUD0-2hbeH51EJ7UqTKTOQANykgAtwtiX13s9-pUi1mVj1aabKkXjM8abc__j9kfivemXbKEfAEOyCtrVANEnOQndwg2fWjCcei4qVSVhlbcknP7BS4u6DJGmYTGFO3MazaBSWrL77mTR67GLQ8QcEmHT1nBHHg5NS2F7HILasJXASW-UW94a7lBG-Ylp932PYruZ-1YO1qvsz8DQNWi41cc5xhmTKhqfrnqMMqc741TRDkZ1ijpGNql9UdenH6BeOvgK0DFhew3xYo1w19BrYAr3XCWIBatrN3A-1_68vVzZBj_wUhON8tSXV_EHp1_RnR12cPnQgvpCIcEmX5q3YMiWtCuB8jwugZKGFEd6EfHeHzLv-4VxuOm_WWZguAuAoZCbSU1V7LiFAaB8-r6MzGu3BrzxYCAO1VSu8jsXZ0sDQgktiaNf6V5O6NrCAORZW7Wvuq_g8wHVStCLFehekNdNElV6-6706kPSFKtHkXOnAmmPhUECIGS0caMY09N5iTGgw6QzrbSywjDw-d-xKzP6c2a1oWr_KY1pP5F5uYQbApn48XPQCQ6nDy2VROVWR9z0_JY-3-cnYtweyx4oT7GzceNsarcWTHjo0IPSc2720is5ttBZU1DY5Jk_ohbWlIw9T5Sxm2xYt7PKnxyA-qDCSN7yYjQbEbZrkeVNy.UbvLKQOP0nfwlx-ubbrfmA")
//	//
//	//res, err := http.DefaultClient.Do(req)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//
//	//defer res.Body.Close()
//	//body, _ := io.ReadAll(res.Body)
//	//
//	//fmt.Println(string(body))
//	//
//	//package main
//	//
//	//import (
//	//	pinecone "github.com/nekomeowww/go-pinecone"
//	//)
//
//	p, err := pinecone.NewIndexClient(
//		pinecone.WithAPIKey("3109588f-dbae-4e4f-84cb-e71ea6e60816"),
//		pinecone.WithEnvironment("gcp-starter"),
//		pinecone.WithProjectName("vsgpt"),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	p.Debug()
//	// Use your own embedding client/library such as OpenAI embedding API
//	API_KEY := "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"
//
//	llm, err := openai.
//	if err != nil {
//		log.Fatal(err)
//	}
//	prompt := "What would be a good company name for a company that makes colorful socks?"
//	completion, err := llm.Call(context.Background(), prompt)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(completion)
//	ctx := context.Background()
//	myVectorEmbedding, _ := llm.CreateEmbedding(ctx, []string{"this is a sentence I want to embed"})
//
//	params := pinecone.UpsertVectorsParams{
//		Vectors: []*pinecone.Vector{
//			{
//				ID:       "YOUR_VECTOR_ID",
//				Values:   myVectorEmbedding[0],
//				Metadata: map[string]any{"foo": "bar"},
//			},
//		},
//	}
//	resp, err := p.UpsertVectors(ctx, params)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("%+v\n", resp)
//}
