/*
Package metadata handles metada for OPTIONS requests.

Metada has all the information about available actions.
Basic example:

    metadata := New("Product detail").SetDescription("Basic information about product")
    metadata.ActionCreate().From(ProductNew{})
    metadata.ActionDelete()
    metadata.ActionRetrieve().From(Product{})

    metadata.Field('status').Choices.Add(1, "Visible").Add(2. "Hidden")
    metadata.Field('result', 'user', 'status').Choices.Add(1, "Visible").Add(2. "Hidden")

    metadata.RemoveActionCreate().RemoveActionDelete()


Example structure:

    {
        "name": "test endpoint",
        "description": "description",
        "actions": {
            "POST": {
                "type": "struct",
                "fields": {
                    "name": {
                        "type": "string",
                        min_length: 10
                    }
                }
            }
        }
    }
*/
package metadata

