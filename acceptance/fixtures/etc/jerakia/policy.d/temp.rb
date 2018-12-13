policy :temp do

  lookup :schema do
    datasource :file, {
      :docroot => '/var/lib/jerakia/schema',
      :format => :json,
      :enable_caching => true,
      :searchpath => ['']
    }
  end
end
